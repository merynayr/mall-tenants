package user

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/repository/user/converter"
	modelRepo "github.com/merynayr/mall-tenants/internal/repository/user/model"
	"github.com/merynayr/mall-tenants/internal/utils/hash"
)

// Константы названий столбцов БД
const (
	usersTable   = "users"
	clientsTable = "clients"

	idColumn   = "id"
	nameColumn = "username"

	UserIDColumn   = "user_id"
	EmailColumn    = "email"
	PasswordColumn = "password_hash"
	RoleColumn     = "role"

	OrganizationNameColumn = "organization_name"
	ContactPersonColumn    = "contact_person"
	AddressColumn          = "address"
	PhoneColumn            = "phone"
	RequisitesColumn       = "requisites"
	ClientIDColumn         = "client_id"
)

const role = 0

// Структура репо с клиентом базы данных (интерфейсом)
type repo struct {
	db db.Client
}

// NewRepository возвращает новый объект репо слоя
func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) CreateUser(ctx context.Context, req *model.RegisterRequest) (int64, error) {
	passHash, err := hash.EncryptPassword(req.Password)
	if err != nil {
		return 0, err
	}

	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	tx, err := r.db.DB().BeginTx(ctx, txOpts)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil {
			return
		}
	}()

	query, args, err := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns(
			EmailColumn,
			PasswordColumn,
			RoleColumn,
		).
		Values(
			req.Email,
			passHash,
			role,
		).
		Suffix("RETURNING user_id").
		ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.CreateUser_UsersTable",
		QueryRaw: query,
	}
	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	query, args, err = sq.Insert(clientsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(
			ClientIDColumn,
			OrganizationNameColumn,
			ContactPersonColumn,
			AddressColumn,
			PhoneColumn,
			RequisitesColumn,
		).
		Values(
			userID,
			req.OrganizationName,
			req.ContactPerson,
			req.Address,
			req.Phone,
			req.Requisites,
		).
		Suffix("RETURNING client_id").
		ToSql()

	if err != nil {
		return 0, err
	}

	q = db.Query{
		Name:     "user_repository.CreateUser_ClientsTable",
		QueryRaw: query,
	}

	err = r.db.DB().ScanOneContext(ctx, &userID, q, args...)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// GetUserByEmail получает из БД информацию пользователя
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*model.User, bool, error) {
	query, args, err := sq.Select(UserIDColumn, EmailColumn, PasswordColumn, RoleColumn).
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{EmailColumn: email}).
		Limit(1).ToSql()

	if err != nil {
		return nil, false, err
	}

	q := db.Query{
		Name:     "user_repository.GetUserByEmail",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return converter.ToUserFromRepo(&user), true, nil
}

// UpdateUser обновляет данные пользователя по id
func (r *repo) UpdateUser(ctx context.Context, user *model.UserUpdate) error {
	builderUpdate := sq.Update(usersTable).
		PlaceholderFormat(sq.Dollar)

	if user.Username != "" {
		builderUpdate = builderUpdate.Set(nameColumn, &user.Username)
	}

	if user.Role >= 0 {
		builderUpdate = builderUpdate.Set(RoleColumn, &user.Role)
	}

	builderUpdate = builderUpdate.Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdateUser",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

// IsEmailExist проверяет, существует ли в БД указанный email
func (r *repo) IsEmailExist(ctx context.Context, email string) (bool, error) {
	query, args, err := sq.Select("1").
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{EmailColumn: email}).
		Limit(1).ToSql()

	if err != nil {
		return false, err
	}

	q := db.Query{
		Name:     "user_repository.IsEmailExist",
		QueryRaw: query,
	}

	var one int

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&one)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
