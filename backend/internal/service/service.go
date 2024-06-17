package service

import (
	"context"
	"main/internal/models"
)

type Repository interface {
	AddGood(ctx context.Context, good models.Good) (err error)
	AddGoodStock(ctx context.Context, goodStock models.GoodStock) (err error)
	AddStock(ctx context.Context, stock models.Stock) (err error)
	AddUser(ctx context.Context, user models.User) (err error)

	DeleteGoodByID(ctx context.Context, good models.Good) (err error)
	DeleteStockByID(ctx context.Context, stock models.Stock) (err error)
	DeleteUserByID(ctx context.Context, user models.User) (err error)

	EditGood(ctx context.Context, good models.Good) (err error)
	EditGoodStock(ctx context.Context, goodStock models.GoodStock) (err error)
	EditStock(ctx context.Context, stock models.Stock) (err error)
	EditUser(ctx context.Context, user models.User) (err error)

	GetAllGoods(ctx context.Context) (goods []models.Good, err error)
	GetAllStocks(ctx context.Context) (stocks []models.Stock, err error)
	GetAllUsers(ctx context.Context) (users []models.User, err error)
	GetUserByLogin(ctx context.Context, login string) (user models.User, err error)

	SearchGoods(ctx context.Context, searchRequest models.SearchGoodRequest) (goods []models.Good, err error)
	SearchUsers(ctx context.Context, searchRequest models.SearchUserRequest) (users []models.User, err error)
	SearchStocks(ctx context.Context, searchRequest models.SearchStockRequest) (stocks []models.Stock, err error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
