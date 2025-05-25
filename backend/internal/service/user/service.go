package user

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/model"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"
	"github.com/merynayr/mall-tenants/internal/sys"
)

// Структура сервисного слоя с объектами репо слоя
// и транзакционного менеджера
type userService struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService возвращает объект сервисного слоя
func NewService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &userService{
		userRepository: userRepository,
		txManager:      txManager,
	}
}

func (s *userService) GetClients(ctx context.Context, limit, offset uint64) ([]model.Client, error) {
	if limit <= 0 {
		limit = 20
	}
	clients, err := s.userRepository.GetClients(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	if len(clients) == 0 {
		return nil, sys.ClientsNotFoundError
	}
	return clients, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, exists, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sys.UserNotFoundError
	}

	return user, nil
}

func (s *userService) CreateClient(ctx context.Context, req model.RegisterRequest) error {
	if req.Password != req.ConfirmPassword {
		return sys.PasswordsDoNotMatchError
	}

	_, exist, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if exist {
		return sys.UserAlreadyExistsError
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error

		userID, errTx := s.userRepository.CreateUser(ctx, &req)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRepository.CreateClient(ctx, &req, userID)
		if errTx != nil {
			return errTx
		}
		return nil
	})

	return err
}
