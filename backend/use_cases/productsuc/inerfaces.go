package productsuc

import "github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/domain/entities"

type ProductsDataStorer interface {
	GetAll() []entities.Product
	Create(product *entities.Product) error
}
