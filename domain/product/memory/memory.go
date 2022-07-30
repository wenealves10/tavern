package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/aggregate"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	// Collect all Products from map
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mrp *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mrp.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

func (mpr *MemoryProductRepository) Update(upprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}