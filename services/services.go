package services

import "tsProject/models"

type Services struct {
	UserService        models.IUserService
	ProductService     models.IProductService
	CycleService       models.ICycleService
	AuthService        models.IAuthService
	InstitutionService models.IInstitutionService
}

func CreateNewServices(
	userRepositories models.IUserRepository,
	productRepositories models.IProductRepository,
	cycleRepositories models.ICycleRepository,
	institutionRepositories models.IInstitutionRepository,
	logger models.ILogger,
) *Services {
	userService := NewUserService(userRepositories, logger)
	productService := NewProductService(productRepositories, logger)
	cycleService := NewCycleService(cycleRepositories, logger)
	authService := NewAuthService(userRepositories, logger)
	institutionService := NewInstitutionService(institutionRepositories, logger)

	return &Services{
		UserService:        userService,
		ProductService:     productService,
		CycleService:       cycleService,
		AuthService:        authService,
		InstitutionService: institutionService,
	}
}
