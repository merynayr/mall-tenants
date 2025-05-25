package sys

import "github.com/merynayr/mall-tenants/internal/sys/codes"

// Текстовые сообщения об ошибках
const (
	// Общие
	ErrInvalidRequest            = "неверный запрос"
	ErrApplicationSubmitFailed   = "не удалось отправить заявку"
	ErrApplicationsNotFound      = "заявки не найдены"
	ErrFloorPlanSaveFailed       = "не удалось сохранить план этажа"
	ErrFloorNumberMustBePositive = "номер этажа должен быть положительным"
	ErrMonthlyRateMustBePositive = "месячная плата должна быть положительной"
	ErrNoPaymentIDsProvided      = "не переданы идентификаторы платежей"
	ErrMarkPaymentsFailed        = "не удалось отметить платежи как оплаченные"
	ErrPremiseNotFound           = "помещение не найдено"
	ErrRentalsNotFound           = "аренды не найдены"
	ErrClientsNotFound           = "клиенты не найдены"
	ErrPaymentsNotFound          = "платежи не найдены"
	ErrFloorPlanNotFound         = "план этажа не найден"
	ErrUserNotFound              = "пользователь не найден"
	ErrUserAlreadyExists         = "пользователь уже существует"
	ErrPremiseOccupied           = "помещение занято"
	ErrPremiseUnderRepair        = "помещение в ремонте и недоступно"
	ErrPremiseAlreadyExists      = "помещение с таким идентификатором уже существует"
	ErrRentalAlreadyExists       = "аренда с таким идентификатором уже существует"
	ErrApplicationAlreadyExists  = "заявка с таким идентификатором уже существует"
	ErrInvalidIDFormat           = "неверный формат идентификатора"
	ErrInternalServer            = "внутренняя ошибка сервера"

	// Авторизация и токены
	ErrInvalidPassword         = "неверный пароль"
	ErrInvalidCredentials      = "неверный логин или пароль"
	ErrPasswordsDoNotMatch     = "пароль и пароль подтверждения не совпадают"
	ErrUserUnauthorized        = "пользователь не авторизован"
	ErrAccessDenied            = "доступ запрещён"
	ErrAuthHeaderMissing       = "отсутствует заголовок авторизации"
	ErrAuthHeaderInvalidFormat = "неверный формат заголовка авторизации"
	ErrInvalidAccessToken      = "недействительный access-токен"
	ErrTokenExpired            = "срок действия токена истёк"
	ErrInvalidRefreshToken     = "недействительный refresh-токен"
	ErrTokenGenerationFailed   = "ошибка при генерации токена"
)

// Готовые объекты ошибок с поясняющими комментариями
var (
	// InvalidRequestError возникает, когда передан некорректный запрос. Код ошибки: 400 (Bad Request)
	InvalidRequestError = NewCommonError(ErrInvalidRequest, codes.BadRequest)

	// ApplicationSubmitError возникает, когда не удаётся отправить заявку. Код ошибки: 500 (InternalServerError Server Error)
	ApplicationSubmitError = NewCommonError(ErrApplicationSubmitFailed, codes.InternalServerError)

	// ApplicationsNotFoundError возникает, когда список заявок пустой. Код ошибки: 404 (Not Found)
	ApplicationsNotFoundError = NewCommonError(ErrApplicationsNotFound, codes.NotFound)

	// FloorPlanSaveError возникает, когда не удаётся сохранить SVG-план этажа. Код ошибки: 500 (InternalServerError Server Error)
	FloorPlanSaveError = NewCommonError(ErrFloorPlanSaveFailed, codes.InternalServerError)

	// FloorNumberInvalidError возникает, когда номер этажа указан некорректно (например, отрицательное число). Код ошибки: 400 (Bad Request)
	FloorNumberInvalidError = NewCommonError(ErrFloorNumberMustBePositive, codes.BadRequest)

	// MonthlyRateInvalidError возникает, когда значение ежемесячной платы указано как неположительное число. Код ошибки: 400 (Bad Request)
	MonthlyRateInvalidError = NewCommonError(ErrMonthlyRateMustBePositive, codes.BadRequest)

	// NoPaymentIDsError возникает, когда не переданы идентификаторы платежей для массовой обработки. Код ошибки: 400 (Bad Request)
	NoPaymentIDsError = NewCommonError(ErrNoPaymentIDsProvided, codes.BadRequest)

	// MarkPaymentsFailedError возникает, когда не удаётся отметить платежи как оплаченные. Код ошибки: 500 (InternalServerError Server Error)
	MarkPaymentsFailedError = NewCommonError(ErrMarkPaymentsFailed, codes.InternalServerError)

	// PremiseNotFoundError возникает, когда запрашиваемое помещение не найдено. Код ошибки: 404 (Not Found)
	PremiseNotFoundError = NewCommonError(ErrPremiseNotFound, codes.NotFound)

	// RentalsNotFoundError возникает, когда аренды не найдены для выбранного помещения или пользователя. Код ошибки: 404 (Not Found)
	RentalsNotFoundError = NewCommonError(ErrRentalsNotFound, codes.NotFound)

	// PaymentsNotFoundError возникает, когда платежи не найдены. Код ошибки: 404 (Not Found)
	PaymentsNotFoundError = NewCommonError(ErrPaymentsNotFound, codes.NotFound)

	// ClientsNotFoundError возникает, когда клиенты не найдены. Код ошибки: 404 (Not Found)
	ClientsNotFoundError = NewCommonError(ErrClientsNotFound, codes.NotFound)

	// FloorPlanNotFoundError возникает, когда план этажа не найден. Код ошибки: 404 (Not Found)
	FloorPlanNotFoundError = NewCommonError(ErrFloorPlanNotFound, codes.NotFound)

	// UserNotFoundError возникает, когда запрашиваемый пользователь не найден. Код ошибки: 404 (Not Found)
	UserNotFoundError = NewCommonError(ErrUserNotFound, codes.NotFound)

	// UserAlreadyExistsError возникает, когда создаваемый пользователь уже существует. Код ошибки: 409 (Conflict)
	UserAlreadyExistsError = NewCommonError(ErrUserAlreadyExists, codes.Conflict)

	// PremiseAlreadyExistsError возникает, когда помещение с таким ID уже существует. Код ошибки: 409 (Conflict)
	PremiseAlreadyExistsError = NewCommonError(ErrPremiseAlreadyExists, codes.Conflict)

	// RentalAlreadyExistsError возникает, когда аренда с таким ID уже существует. Код ошибки: 409 (Conflict)
	RentalAlreadyExistsError = NewCommonError(ErrRentalAlreadyExists, codes.Conflict)

	// ApplicationAlreadyExistsError возникает, когда заявка с таким ID уже существует. Код ошибки: 409 (Conflict)
	ApplicationAlreadyExistsError = NewCommonError(ErrApplicationAlreadyExists, codes.Conflict)

	// PremiseOccupiedError возникает, когда помещение уже занято. Код ошибки: 409 (Conflict)
	PremiseOccupiedError = NewCommonError(ErrPremiseOccupied, codes.Conflict)

	// PremiseUnderRepairError возникает, когда помещение в ремонте и недоступно. Код ошибки: 403 (Forbidden)
	PremiseUnderRepairError = NewCommonError(ErrPremiseUnderRepair, codes.Forbidden)

	// InvalidIDFormatError возникает, когда передан идентификатор в неверном формате (например, некорректный UUID). Код ошибки: 400 (Bad Request)
	InvalidIDFormatError = NewCommonError(ErrInvalidIDFormat, codes.BadRequest)

	// InternalServerErrorServerError универсальная ошибка для неожиданных сбоев на сервере. Код ошибки: 500 (InternalServerError Server Error)
	InternalServerError = NewCommonError(ErrInternalServer, codes.InternalServerError)

	// InvalidPasswordError возникает, когда введён неверный пароль. Код ошибки: 403 (Forbidden)
	InvalidPasswordError = NewCommonError(ErrInvalidPassword, codes.Forbidden)

	// InvalidCredentialsError возникает, когда введён неверный логин или пароль. Код ошибки: 403 (Forbidden)
	InvalidCredentialsError = NewCommonError(ErrInvalidCredentials, codes.Forbidden)

	// PasswordsDoNotMatchError возникает, когда пароль и повтор пароля не сопадают. Код ошибки: 400 (Bad Request)
	PasswordsDoNotMatchError = NewCommonError(ErrPasswordsDoNotMatch, codes.BadRequest)

	// UserUnauthorizedError возникает, когда пользователь не авторизован. Код ошибки: 401 (Unauthorized)
	UserUnauthorizedError = NewCommonError(ErrUserUnauthorized, codes.Unauthorized)

	// AccessDeniedError возникает, когда у пользователя нет доступа к ресурсу. Код ошибки: 403 (Forbidden)
	AccessDeniedError = NewCommonError(ErrAccessDenied, codes.Forbidden)

	// AuthHeaderMissingError возникает, когда отсутствует заголовок авторизации. Код ошибки: 401 (Unauthorized)
	AuthHeaderMissingError = NewCommonError(ErrAuthHeaderMissing, codes.Unauthorized)

	// AuthHeaderInvalidFormatError возникает, когда формат заголовка авторизации некорректен. Код ошибки: 401 (Unauthorized)
	AuthHeaderInvalidFormatError = NewCommonError(ErrAuthHeaderInvalidFormat, codes.Unauthorized)

	// InvalidAccessTokenError возникает, когда передан недействительный access-токен. Код ошибки: 401 (Unauthorized)
	InvalidAccessTokenError = NewCommonError(ErrInvalidAccessToken, codes.Unauthorized)

	// TokenExpiredError возникает, когда срок действия access-токена истёк. Код ошибки: 401 (Unauthorized)
	TokenExpiredError = NewCommonError(ErrTokenExpired, codes.Unauthorized)

	// InvalidRefreshTokenError возникает, когда refresh-токен недействителен. Код ошибки: 403 (Forbidden)
	InvalidRefreshTokenError = NewCommonError(ErrInvalidRefreshToken, codes.Forbidden)

	// TokenGenerationFailedError возникает при ошибке генерации access или refresh токенов. Код ошибки: 500 (InternalServerError Server Error)
	TokenGenerationFailedError = NewCommonError(ErrTokenGenerationFailed, codes.InternalServerError)
)
