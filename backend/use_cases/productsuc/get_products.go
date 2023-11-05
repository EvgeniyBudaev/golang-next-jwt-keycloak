package productsuc

import (
	"context"
	"github.com/EvgeniyBudaev/golang-next-jwt-keycloak/backend/domain/entities"
)

type getProductsUseCase struct {
	dataStore ProductsDataStorer
}

func NewGetProductsUseCase(ds ProductsDataStorer) *getProductsUseCase {
	return &getProductsUseCase{
		dataStore: ds,
	}
}

func (uc *getProductsUseCase) GetProducts(ctx context.Context) []entities.Product {
	all := uc.dataStore.GetAll()
	return all
}
