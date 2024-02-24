package lru

import "time"

type LRU struct {
	head         *Node
	tail         *Node
	length       int
	maxLength    int
	evictionTime int
}

type Node struct {
	data Item
	next *Node
	prev *Node
}

type Item struct {
	key        string
	value      string
	expiration time.Time
}

type SetRequest struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}
