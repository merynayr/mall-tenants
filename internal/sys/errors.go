package sys

import "github.com/merynayr/mall-tenants/internal/sys/codes"

// Константы с текстами ошибок
const (
	ErrNotFound                = "not found"
	ErrNotEnoughCoins          = "not enough coins"
	ErrSelfTransferNotAllowed  = "you can't transfer money to yourself"
	ErrInvalidRefreshToken     = "invalid refresh token"
	ErrInvalidPassword         = "invalid password"
	ErrAuthHeaderNotProvided   = "authorization header is not provided"
	ErrInvalidAuthHeaderFormat = "invalid authorization header format"
	ErrInvalidAccessToken      = "access token is invalid"
	ErrAccessDenied            = "access denied"
	ErrInvalidRequest          = "invalid request"
	ErrInvalidUser             = "invalid user"
	ErrUserNotFound            = "user not found"
	ErrUserExist               = "user already exist"
	ErrRecipientNotFound       = "recipient not found"
	ErrPasswordsDoNotMatch     = "password and confirm password do not match"
	ErrExist                   = "already exist"
	ErrPaymentNotFound         = "payment not found"
	ErrInvalidPaymentAmount    = "invalid payment amount"
	ErrPaymentAlreadyExists    = "payment already exists"
	ErrRentalNotFound          = "rental not found"
	ErrOverduePaymentsNotFound = "overdue payments not found"
	ErrPremiseNotFound         = "premise not found"
	ErrPremiseOccupied         = "premise is already occupied"
	ErrPremiseUnderRepair      = "premise is under repair and unavailable"
)

// Готовые объекты ошибок с комментариями
var (
	// ItemNotFoundError возникает, когда запрашиваемый объект не найден. Код ошибки: 404 (Not Found)
	NotFoundError = NewCommonError(ErrNotFound, codes.NotFound)

	// NotEnoughCoinsError возникает, когда у пользователя недостаточно монет для перевода. Код ошибки: 400 (Bad Request)
	NotEnoughCoinsError = NewCommonError(ErrNotEnoughCoins, codes.BadRequest)

	// SelfTransferNotAllowedError возникает, когда пользователь пытается перевести монеты самому себе. Код ошибки: 400 (Bad Request)
	SelfTransferNotAllowedError = NewCommonError(ErrSelfTransferNotAllowed, codes.BadRequest)

	// InvalidRefreshTokenError возникает, когда передан недействительный refresh-токен. Код ошибки: 401 (Unauthorized)
	InvalidRefreshTokenError = NewCommonError(ErrInvalidRefreshToken, codes.Unauthorized)

	// InvalidPasswordError возникает, когда введён неверный пароль. Код ошибки: 401 (Unauthorized)
	InvalidPasswordError = NewCommonError(ErrInvalidPassword, codes.Unauthorized)

	// AuthHeaderNotProvidedError возникает, когда отсутствует заголовок авторизации. Код ошибки: 401 (Unauthorized)
	AuthHeaderNotProvidedError = NewCommonError(ErrAuthHeaderNotProvided, codes.Unauthorized)

	// InvalidAuthHeaderFormatError возникает, когда формат заголовка авторизации неверный. Код ошибки: 401 (Unauthorized)
	InvalidAuthHeaderFormatError = NewCommonError(ErrInvalidAuthHeaderFormat, codes.Unauthorized)

	// InvalidAccessTokenError возникает, когда передан недействительный access-токен. Код ошибки: 401 (Unauthorized)
	InvalidAccessTokenError = NewCommonError(ErrInvalidAccessToken, codes.Unauthorized)

	// AccessDeniedError возникает, когда у пользователя недостаточно прав для выполнения операции. Код ошибки: 403 (Forbidden)
	AccessDeniedError = NewCommonError(ErrAccessDenied, codes.Forbidden)

	// InvalidRequestError возникает, когда передан некорректный запрос. Код ошибки: 400 (Bad Request)
	InvalidRequestError = NewCommonError(ErrInvalidRequest, codes.BadRequest)

	// InvalidUserError возникает, когда передан недействительный пользователь. Код ошибки: 400 (Bad Request)
	InvalidUserError = NewCommonError(ErrInvalidUser, codes.BadRequest)

	// UserNotFoundError возникает, когда запрашиваемый пользователь не найден. Код ошибки: 404 (Not Found)
	UserNotFoundError = NewCommonError(ErrUserNotFound, codes.NotFound)

	// UserExistError возникает, когда регистрируемый пользователь уже существует. Код ошибки: 409 (Conflict)
	UserExistError = NewCommonError(ErrUserExist, codes.Conflict)

	// RecipientNotFoundError возникает, когда запрашиваемый получатель перевода не найден. Код ошибки: 404 (Not Found)
	RecipientNotFoundError = NewCommonError(ErrRecipientNotFound, codes.NotFound)

	// PasswordsDoNotMatchError возникает, когда пароль и повтор пароля не сопадают. Код ошибки: 400 (Bad Request)
	PasswordsDoNotMatchError = NewCommonError(ErrPasswordsDoNotMatch, codes.BadRequest)

	// ExistError возникает, когда создаваемый объект уже существует. Код ошибки: 409 (Conflict)
	ExistError = NewCommonError(ErrExist, codes.Conflict)

	// PaymentNotFoundError возникает, когда запрашиваемый платеж не найден. Код ошибки: 404 (Not Found)
	PaymentNotFoundError = NewCommonError(ErrPaymentNotFound, codes.NotFound)

	// InvalidPaymentAmountError возникает, когда сумма платежа некорректна. Код ошибки: 400 (Bad Request)
	InvalidPaymentAmountError = NewCommonError(ErrInvalidPaymentAmount, codes.BadRequest)

	// PaymentAlreadyExistsError возникает, когда платеж уже существует. Код ошибки: 409 (Conflict)
	PaymentAlreadyExistsError = NewCommonError(ErrPaymentAlreadyExists, codes.Conflict)

	// RentalNotFoundError возникает, когда запрашиваемая аренда не найдена. Код ошибки: 404 (Not Found)
	RentalNotFoundError = NewCommonError(ErrRentalNotFound, codes.NotFound)

	// OverduePaymentsNotFoundError возникает, когда просроченные платежи не найдены. Код ошибки: 404 (Not Found)
	OverduePaymentsNotFoundError = NewCommonError(ErrOverduePaymentsNotFound, codes.NotFound)

	// PremiseNotFoundError возникает, когда запрашиваемое помещение не найдено. Код ошибки: 404 (Not Found)
	PremiseNotFoundError = NewCommonError(ErrPremiseNotFound, codes.NotFound)

	// PremiseOccupiedError возникает, когда помещение уже занято. Код ошибки: 409 (Conflict)
	PremiseOccupiedError = NewCommonError(ErrPremiseOccupied, codes.Conflict)

	// PremiseUnderRepairError возникает, когда помещение в ремонте и недоступно. Код ошибки: 403 (Forbidden)
	PremiseUnderRepairError = NewCommonError(ErrPremiseUnderRepair, codes.Forbidden)
)
