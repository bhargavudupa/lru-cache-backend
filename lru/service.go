package lru

import (
	"sync"
	"time"
)

type LruService interface {
	Get(key string) (value string, err error)
	Set(key string, value string)
}

type lruService struct {
	lru LRU
}

func NewLruService() LruService {
	lruService := lruService{}
	lruService.lru = *NewLRU(1024, 60)

	var mu sync.Mutex
	go func() {
		for {
			mu.Lock()
			lruService.lru.Evict()
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	return &lruService
}

func (service *lruService) Get(key string) (string, error) {
	val, err := service.lru.Get(key)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (service *lruService) Set(key string, value string) {
	service.lru.Set(key, value)
}
