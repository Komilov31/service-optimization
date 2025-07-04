package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/Komilov31/l0/internal/model"
	"github.com/google/uuid"
)

type InMemoryCache struct {
	mu           sync.RWMutex
	cacheStorage map[uuid.UUID]model.Order
	repo         model.Storage
}

func New(repo model.Storage) *InMemoryCache {
	return &InMemoryCache{
		mu:           sync.RWMutex{},
		cacheStorage: make(map[uuid.UUID]model.Order),
		repo:         repo,
	}
}

func (c *InMemoryCache) LoadFromDbToCache() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	lastOrders := []model.Order{}

	lastOrders, err := c.repo.GetLastOrders(context.Background())
	if err != nil {
		return fmt.Errorf("could not load orders from db: %w", err)
	}

	for _, order := range lastOrders {
		c.cacheStorage[order.OrderUid] = order
	}

	return nil
}

func (c *InMemoryCache) IsOrderInCache(orderUid uuid.UUID) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.cacheStorage[orderUid]
	return ok
}

func (c *InMemoryCache) SaveToCache(order model.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheStorage[order.OrderUid] = order
}

func (c *InMemoryCache) GetOrderById(orderUid uuid.UUID) model.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.cacheStorage[orderUid]
}
