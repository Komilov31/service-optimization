package cache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Komilov31/l0/internal/model"
	"github.com/google/uuid"
)

type InMemoryCache struct {
	mu           sync.RWMutex
	cacheStorage map[uuid.UUID]cacheItem
	repo         model.Storage
}

type cacheItem struct {
	order      model.Order
	expiration time.Time
}

func New(repo model.Storage) *InMemoryCache {
	return &InMemoryCache{
		mu:           sync.RWMutex{},
		cacheStorage: make(map[uuid.UUID]cacheItem),
		repo:         repo,
	}
}

func (c *InMemoryCache) LoadFromDbToCache() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var lastOrders []model.Order

	lastOrders, err := c.repo.GetLastOrders(context.Background())
	if err != nil {
		return fmt.Errorf("could not load orders from db: %w", err)
	}

	for _, order := range lastOrders {
		item := cacheItem{
			order:      order,
			expiration: time.Now().Add(time.Hour),
		}
		c.cacheStorage[order.OrderUid] = item
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

	item := cacheItem{
		order:      order,
		expiration: time.Now().Add(time.Hour),
	}
	c.cacheStorage[order.OrderUid] = item
}

func (c *InMemoryCache) RemoveFromCache(orderUid uuid.UUID) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.cacheStorage, orderUid)
}

func (c *InMemoryCache) GetOrderById(orderUid uuid.UUID) model.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.cacheStorage[orderUid].order
}

func (c *InMemoryCache) StartCleaner() {
	go func() {
		ticker := time.NewTicker(time.Hour)
		for range ticker.C {
			c.scanAndClean()
		}
	}()
}

func (c *InMemoryCache) scanAndClean() {
	for key, item := range c.cacheStorage {
		if time.Until(item.expiration) <= 0 {
			c.RemoveFromCache(key)
		}
	}
}
