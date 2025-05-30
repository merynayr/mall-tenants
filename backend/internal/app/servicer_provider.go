package app

import (
	"context"
	"log"

	"github.com/merynayr/mall-tenants/internal/client/db"
	"github.com/merynayr/mall-tenants/internal/client/db/pg"
	"github.com/merynayr/mall-tenants/internal/client/db/transaction"
	"github.com/merynayr/mall-tenants/internal/closer"
	"github.com/merynayr/mall-tenants/internal/config"
	"github.com/merynayr/mall-tenants/internal/config/env"
	"github.com/merynayr/mall-tenants/internal/repository"
	"github.com/merynayr/mall-tenants/internal/service"

	"github.com/merynayr/mall-tenants/internal/api/application"
	"github.com/merynayr/mall-tenants/internal/api/auth"
	"github.com/merynayr/mall-tenants/internal/api/contract"
	"github.com/merynayr/mall-tenants/internal/api/payment"
	"github.com/merynayr/mall-tenants/internal/api/premise"
	"github.com/merynayr/mall-tenants/internal/api/rental"
	"github.com/merynayr/mall-tenants/internal/api/user"

	floorPlan "github.com/merynayr/mall-tenants/internal/api/floorPlan"
	floorPlanRepository "github.com/merynayr/mall-tenants/internal/repository/floorPlan"
	floorPlanService "github.com/merynayr/mall-tenants/internal/service/floorPlan"

	accessService "github.com/merynayr/mall-tenants/internal/service/access"
	authService "github.com/merynayr/mall-tenants/internal/service/auth"

	userRepository "github.com/merynayr/mall-tenants/internal/repository/user"
	userService "github.com/merynayr/mall-tenants/internal/service/user"

	premiseRepository "github.com/merynayr/mall-tenants/internal/repository/premises"
	premiseService "github.com/merynayr/mall-tenants/internal/service/premises"

	contractRepository "github.com/merynayr/mall-tenants/internal/repository/contracts"
	contractService "github.com/merynayr/mall-tenants/internal/service/contracts"

	rentalRepository "github.com/merynayr/mall-tenants/internal/repository/rentals"
	rentalService "github.com/merynayr/mall-tenants/internal/service/rentals"

	paymentRepository "github.com/merynayr/mall-tenants/internal/repository/payments"
	paymentService "github.com/merynayr/mall-tenants/internal/service/payments"

	applicationRepository "github.com/merynayr/mall-tenants/internal/repository/applications"
	applicationService "github.com/merynayr/mall-tenants/internal/service/applications"

	"github.com/merynayr/mall-tenants/internal/middleware"
)

// Структура приложения со всеми зависимости
type serviceProvider struct {
	pgConfig      config.PGConfig
	httpConfig    config.HTTPConfig
	loggerConfig  config.LoggerConfig
	swaggerConfig config.SwaggerConfig
	authConfig    config.AuthConfig
	accessConfig  config.AccessConfig

	dbClient  db.Client
	txManager db.TxManager

	userAPI        *user.API
	userService    service.UserService
	userRepository repository.UserRepository

	authAPI     *auth.API
	authService service.AuthService

	mallAPI           *premise.API
	premiseService    service.PremiseService
	premiseRepository repository.PremiseRepository

	floorPlanAPI        *floorPlan.API
	floorPlanService    service.FloorPlanService
	floorPlanRepository repository.FloorPlanRepository

	rentalAPI        *rental.API
	rentalService    service.RentalService
	rentalRepository repository.RentalRepository

	contractAPI        *contract.API
	contractService    service.ContractService
	contractRepository repository.ContractRepository

	paymentAPI        *payment.API
	paymentService    service.PaymentService
	paymentRepository repository.PaymentRepository

	applicationAPI        *application.API
	applicationService    service.ApplicationService
	applicationRepository repository.ApplicationRepository

	middleware    middleware.Middleware
	accessService service.AccessService
}

// NewServiceProvider возвращает новый объект API слоя
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config:%v", err)
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

// AuthConfig инициализирует конфиг auth сервиса
func (s *serviceProvider) AuthConfig() config.AuthConfig {
	if s.authConfig == nil {
		cfg, err := env.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get auth config")
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

// AccessConfig инициализирует конфиг access конфига
func (s *serviceProvider) AccessConfig() config.AccessConfig {
	if s.accessConfig == nil {
		cfg, err := env.NewAccessConfig()
		if err != nil {
			log.Fatalf("failed to get access service")
		}

		s.accessConfig = cfg
	}

	return s.accessConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("db ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// UserAPI инициализирует api слой user
func (s *serviceProvider) UserAPI(ctx context.Context) *user.API {
	if s.userAPI == nil {
		s.userAPI = user.NewAPI(s.UserService(ctx))
	}

	return s.userAPI
}

// UserService иницилизирует сервисный слой auth
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// AuthAPI инициализирует api слой auth
func (s *serviceProvider) AuthAPI(ctx context.Context) *auth.API {
	if s.authAPI == nil {
		s.authAPI = auth.NewAPI(s.AuthService(ctx), s.AuthConfig())
	}

	return s.authAPI
}

// AuthService иницилизирует сервисный слой auth
func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserRepository(ctx),
			s.AuthConfig(),
		)
	}

	return s.authService
}

// Middleware инициализирует middleware доступа
func (s *serviceProvider) Middleware(ctx context.Context) middleware.Middleware {
	if s.middleware == nil {
		s.middleware = middleware.NewMiddlewareProvider(
			s.AccessService(ctx),
			s.AuthConfig(),
		)
	}
	return s.middleware
}

// AccessService иницилизирует сервисный слой access
func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		uMap, err := s.AccessConfig().UserAccessesMap()
		if err != nil {
			log.Fatalf("failed to get user access map: %v", err)
		}

		s.accessService = accessService.NewService(s.UserService(ctx), uMap, s.AuthConfig())
	}

	return s.accessService
}

// MallAPI инициализирует api слой Mall
func (s *serviceProvider) MallAPI(ctx context.Context) *premise.API {
	if s.mallAPI == nil {
		s.mallAPI = premise.NewAPI(s.PremiseService(ctx))
	}

	return s.mallAPI
}

// PremiseService иницилизирует сервисный слой для помещений
func (s *serviceProvider) PremiseService(ctx context.Context) service.PremiseService {
	if s.premiseService == nil {
		s.premiseService = premiseService.NewService(
			s.PremiseRepository(ctx),
		)
	}

	return s.premiseService
}

func (s *serviceProvider) PremiseRepository(ctx context.Context) repository.PremiseRepository {
	if s.premiseRepository == nil {
		s.premiseRepository = premiseRepository.NewRepository(s.DBClient(ctx))
	}

	return s.premiseRepository
}

// FloorPlanAPI инициализирует api слой Floor Plan
func (s *serviceProvider) FloorPlanAPI(ctx context.Context) *floorPlan.API {
	if s.floorPlanAPI == nil {
		s.floorPlanAPI = floorPlan.NewAPI(s.FloorPlanService(ctx))
	}

	return s.floorPlanAPI
}

// FloorPlanService иницилизирует сервисный слой для помещений
func (s *serviceProvider) FloorPlanService(ctx context.Context) service.FloorPlanService {
	if s.floorPlanService == nil {
		s.floorPlanService = floorPlanService.NewService(
			s.FloorPlanRepository(ctx),
		)
	}

	return s.floorPlanService
}

func (s *serviceProvider) FloorPlanRepository(ctx context.Context) repository.FloorPlanRepository {
	if s.floorPlanRepository == nil {
		s.floorPlanRepository = floorPlanRepository.NewRepository(s.DBClient(ctx))
	}

	return s.floorPlanRepository
}

// RentalAPI инициализирует API-слой для аренды
func (s *serviceProvider) RentalAPI(ctx context.Context) *rental.API {
	if s.rentalAPI == nil {
		s.rentalAPI = rental.NewAPI(s.RentalService(ctx))
	}
	return s.rentalAPI
}

// RentalService инициализирует сервисный слой для аренды
func (s *serviceProvider) RentalService(ctx context.Context) service.RentalService {
	if s.rentalService == nil {
		s.rentalService = rentalService.NewService(
			s.RentalRepository(ctx),
			s.PremiseRepository(ctx),
			s.PaymentRepository(ctx),
			s.ContractRepository(ctx),
			s.TxManager(ctx),
		)
	}
	return s.rentalService
}

// RentalRepository инициализирует репозиторий для аренды
func (s *serviceProvider) RentalRepository(ctx context.Context) repository.RentalRepository {
	if s.rentalRepository == nil {
		s.rentalRepository = rentalRepository.NewRepository(s.DBClient(ctx))
	}
	return s.rentalRepository
}

// PaymentAPI инициализирует API-слой для платежей
func (s *serviceProvider) PaymentAPI(ctx context.Context) *payment.API {
	if s.paymentAPI == nil {
		s.paymentAPI = payment.NewAPI(s.PaymentService(ctx))
	}
	return s.paymentAPI
}

// PaymentService инициализирует сервисный слой для платежей
func (s *serviceProvider) PaymentService(ctx context.Context) service.PaymentService {
	if s.paymentService == nil {
		s.paymentService = paymentService.NewService(
			s.PaymentRepository(ctx),
			s.RentalRepository(ctx),
			s.PremiseRepository(ctx),
			s.TxManager(ctx),
		)
	}
	return s.paymentService
}

// PaymentRepository инициализирует репозиторий для платежей
func (s *serviceProvider) PaymentRepository(ctx context.Context) repository.PaymentRepository {
	if s.paymentRepository == nil {
		s.paymentRepository = paymentRepository.NewRepository(s.DBClient(ctx))
	}
	return s.paymentRepository
}

// ApplicationAPI инициализирует API-слой для заявок
func (s *serviceProvider) ApplicationAPI(ctx context.Context) *application.API {
	if s.applicationAPI == nil {
		s.applicationAPI = application.NewAPI(s.ApplicationService(ctx))
	}
	return s.applicationAPI
}

// ApplicationService инициализирует сервисный слой для заявок
func (s *serviceProvider) ApplicationService(ctx context.Context) service.ApplicationService {
	if s.applicationService == nil {
		s.applicationService = applicationService.NewService(
			s.ApplicationRepository(ctx),
		)
	}
	return s.applicationService
}

// ApplicationRepository инициализирует репозиторий для заявок
func (s *serviceProvider) ApplicationRepository(ctx context.Context) repository.ApplicationRepository {
	if s.applicationRepository == nil {
		s.applicationRepository = applicationRepository.NewRepository(s.DBClient(ctx))
	}
	return s.applicationRepository
}

// ContractAPI инициализирует API-слой для договорв
func (s *serviceProvider) ContractAPI(ctx context.Context) *contract.API {
	if s.contractAPI == nil {
		s.contractAPI = contract.NewAPI(s.ContractService(ctx))
	}
	return s.contractAPI
}

// ContractService инициализирует сервисный слой для договоров
func (s *serviceProvider) ContractService(ctx context.Context) service.ContractService {
	if s.contractService == nil {
		s.contractService = contractService.NewService(
			s.RentalRepository(ctx),
			s.PremiseRepository(ctx),
			s.PaymentRepository(ctx),
			s.ContractRepository(ctx),
			s.TxManager(ctx))
	}
	return s.contractService
}

// ContractRepository инициализирует репозиторий для договоров
func (s *serviceProvider) ContractRepository(ctx context.Context) repository.ContractRepository {
	if s.contractRepository == nil {
		s.contractRepository = contractRepository.NewRepository(s.DBClient(ctx))
	}
	return s.contractRepository
}
