package service

import (
	"context"

	"github.com/Komilov31/l0/internal/model"
	"github.com/google/uuid"
)

type Service struct {
	repo  model.Storage
	cache model.Cache
}

func New(repo model.Storage, cache model.Cache) *Service {
	return &Service{repo: repo, cache: cache}
}

func (s *Service) GetOrderById(ctx context.Context, orderUid uuid.UUID) (model.Order, error) {
	var order model.Order
	if s.cache.IsOrderInCache(orderUid) {
		order = s.cache.GetOrderById(orderUid)
		return order, nil
	}

	order, err := s.repo.GetOrderById(ctx, orderUid)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (s *Service) GetLastOrders(ctx context.Context) ([]model.Order, error) {
	orders, err := s.repo.GetLastOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *Service) CreateOrder(ctx context.Context, order model.Order) error {
	order.OrderUid = uuid.New()
	order.Delivery.DeliveryUid = uuid.New()
	order.Payment.PaymentUid = uuid.New()

	for i := 0; i < len(order.Items); i++ {
		order.Items[i].ItemUid = uuid.New()
	}

	s.cache.SaveToCache(order)

	err := s.repo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) SaveToCache(order model.Order) {
	s.cache.SaveToCache(order)
}
