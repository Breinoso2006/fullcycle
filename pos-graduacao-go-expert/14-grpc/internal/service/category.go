package service

import (
	"context"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/14-grpc/internal/database"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/14-grpc/internal/pb"
)

type CategoryService struct {
	// Utilizado para gerar compatibilidade com outras versões do gRPC.
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse :=
		&pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

	return categoryResponse, nil
}
