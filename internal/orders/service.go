package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/Nipun2001M/go-backend-ecommerce/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound=errors.New("product not found")
	ProductNoStock=errors.New("order stock is greater than available stock")
)
type Service interface {
	PlaceOrder(ctx context.Context,tempOrder createOrderParams) (repo.Order,error)
}

type svc struct {
	repo repo.Queries
	db *pgx.Conn
}

func NewService(repo repo.Queries,db *pgx.Conn) Service{
	return &svc{
		repo: repo,
		db:db,

	}
}


func (s *svc)PlaceOrder(ctx context.Context,tempOrder createOrderParams) (repo.Order,error){
	if tempOrder.CustomerID==0{
		return repo.Order{},fmt.Errorf("customer id is required")
	}
	if len(tempOrder.Items)==0{
		return repo.Order{},fmt.Errorf("at least one item is required")
	}

	tx,err:=s.db.Begin(ctx)
	if err!=nil{
		return repo.Order{},err
	}

	defer tx.Rollback(ctx)

	qtx:=s.repo.WithTx(tx)
	order,err:=qtx.CreateOrder(ctx,tempOrder.CustomerID)
	if err!=nil{
		return repo.Order{},err
	}

	for _,item :=range tempOrder.Items{
		product,err:=s.repo.FindProductById(ctx,item.ProductID)
		if err!=nil{
			return repo.Order{},ErrProductNotFound
		}

		if product.Quantity<item.Quantity{
			return repo.Order{},ProductNoStock
		}

		_,err=qtx.CreateOrderItem(ctx,repo.CreateOrderItemParams{
			OrderID: order.ID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			PriceCents: product.PriceInCenters,		
		})

		if err!=nil{
			return repo.Order{},err}
	
}
tx.Commit(ctx)
return  order,nil
}