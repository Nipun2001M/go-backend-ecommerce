package products

import (
	"context"

	repo "github.com/Nipun2001M/go-backend-ecommerce/internal/adapters/postgresql/sqlc"
)


type Service interface{
	ListProducts(ctx context.Context) ([]repo.Product,error)
	GetProductById(ctx context.Context,id int)(repo.Product,error)


} 

type svc struct{
	repo repo.Querier


}

func NewService(repo repo.Querier)Service{
	return &svc{
		repo: repo,
	}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product,error){
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProductById(ctx context.Context,id int) (repo.Product,error){
	return s.repo.FindProductById(ctx,int64(id))
}

