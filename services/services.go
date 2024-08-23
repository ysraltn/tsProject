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
) *Services {
	userService := NewUserService(userRepositories)
	productService := NewProductService(productRepositories)
	cycleService := NewCycleService(cycleRepositories)
	authService := NewAuthService(userRepositories)
	institutionService := NewInstitutionService(institutionRepositories)

	return &Services{
		UserService:        userService,
		ProductService:     productService,
		CycleService:       cycleService,
		AuthService:        authService,
		InstitutionService: institutionService,
	}
}
