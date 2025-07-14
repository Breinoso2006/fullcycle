//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/19-di/product"
	"github.com/google/wire"
)

var setRepositoryDependencies = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	// defino todos os providers que serão usados para injetar as dependências
	wire.Build(
		setRepositoryDependencies,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
