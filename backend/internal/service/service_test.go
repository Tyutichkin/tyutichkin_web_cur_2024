package service

import (
	"context"
	"main/internal/models"
)

// создаем мок-репозиторий (упрощённый)
type mockRepo struct{}

func (m *mockRepo) AddUser(ctx context.Context, user models.User) error {
	return nil
}

func (m *mockRepo) AddGood(ctx context.Context, good models.Good) (err error) {
	return nil
}
func (m *mockRepo) AddGoodStock(ctx context.Context, goodStock models.GoodStock) (err error) {
	return nil
}
func (m *mockRepo) AddStock(ctx context.Context, stock models.Stock) (err error) {
	return nil
}
func (m *mockRepo) DeleteGoodByID(ctx context.Context, good models.Good) (err error) {
	return nil
}
func (m *mockRepo) DeleteStockByID(ctx context.Context, stock models.Stock) (err error) {
	return nil
}
func (m *mockRepo) DeleteUserByID(ctx context.Context, user models.User) (err error) {
	return nil
}

func (m *mockRepo) EditGood(ctx context.Context, good models.Good) (err error) {
	return nil
}
func (m *mockRepo) EditGoodStock(ctx context.Context, goodStock models.GoodStock) (err error) {
	return nil
}
func (m *mockRepo) EditStock(ctx context.Context, stock models.Stock) (err error) {
	return nil
}
func (m *mockRepo) EditUser(ctx context.Context, user models.User) (err error) {
	return nil
}

func (m *mockRepo) GetAllGoods(ctx context.Context) (goods []models.Good, err error) {
	return nil, nil
}
func (m *mockRepo) GetAllStocks(ctx context.Context) (stocks []models.Stock, err error) {
	return nil, nil
}
func (m *mockRepo) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	return nil, nil
}
func (m *mockRepo) GetUserByLogin(ctx context.Context, login string) (user models.User, err error) {
	return models.User{}, nil
}

func (m *mockRepo) SearchGoods(ctx context.Context, searchRequest models.SearchGoodRequest) (goods []models.Good, err error) {
	return nil, nil
}
func (m *mockRepo) SearchUsers(ctx context.Context, searchRequest models.SearchUserRequest) (users []models.User, err error) {
	return nil, nil
}
func (m *mockRepo) SearchStocks(ctx context.Context, searchRequest models.SearchStockRequest) (stocks []models.Stock, err error) {
	return nil, nil
}
