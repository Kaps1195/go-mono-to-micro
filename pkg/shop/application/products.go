package application

import (
	"errors"

	"github.com/kaps1195/go-mono-to-micro/pkg/common/price"
	"github.com/kaps1195/go-mono-to-micro/pkg/shop/domain/products"
)

type productReadModel interface {
	AllProducts() ([]products.Product, err)
}

type ProductsService struct {
	repo      products.Repository
	readModel productReadModel
}

type AddProductCommand struct {
	ID            string
	Name          string
	Description   string
	Pricecents    uint
	PriceCurrency string
}

func NewProductsService() ProductsService {

}

func (s ProductsService) AllProducts() {

}

func (s ProductsService) AddProduct(cmd AddProductCommand) error {

	price, err := price.NewPrice(cmd.Pricecents, cmd.PriceCurrency)
	if err != nil {
		return errors.Wrap(err, "Invalid product price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, cmd.Description, price)
	if err != nil {
		return errors.Wrap(err, "Cannot create product")
	}

	if err := s.repo.Save(p); err != nil {
		return errors.Wrap(err, "Cannot save product")
	}

	return nil
}
