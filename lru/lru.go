package lru

import (
	"errors"
	"fmt"
	"time"
)

func NewLRU(maxLength int, evictionTime int) *LRU {
	return &LRU{maxLength: maxLength, evictionTime: evictionTime}
}

func (l *LRU) Set(key string, value string) {
	expTime := time.Now().Add(time.Second * time.Duration(l.evictionTime))
	if l.length == 0 {
		newNode := Node{data: Item{key: key, value: value, expiration: expTime}, next: l.head}
		l.head = &newNode
		l.tail = &newNode
		l.length++
		return
	}

	l.Get(key)

	if l.head.data.key == key {
		l.head.data.value = value
		l.head.data.expiration = expTime
	} else {
		if l.length == l.maxLength {
			curNode := l.head
			for i := 0; i < l.length-2; i++ {
				curNode = curNode.next
			}
			l.tail = curNode
			curNode.next = nil
			l.length--
		}
		newNode := Node{data: Item{key: key, value: value, expiration: expTime}, next: l.head}
		newNode.next = l.head
		l.head.prev = &newNode
		l.head = &newNode
		l.length++
	}
}

func (l *LRU) Get(key string) (string, error) {
	if l.length == 0 {
		return "", errors.New("Key not found")
	}
	if l.head.data.key == key {
		return l.head.data.value, nil
	}
	curNode := l.head
	for i := 0; i < l.length-1; i++ {
		if curNode.data.key == key {
			prev := curNode.prev
			next := curNode.next
			prev.next = next
			next.prev = prev
			l.head.prev = curNode
			curNode.next = l.head
			curNode.data.expiration = time.Now().Add(time.Second * time.Duration(l.evictionTime))
			curNode.prev = nil
			l.head = curNode
			return l.head.data.value, nil
		}
		curNode = curNode.next
	}
	if curNode.data.key == key {
		prev := curNode.prev
		prev.next = nil
		l.tail = prev
		l.head.prev = curNode
		curNode.next = l.head
		curNode.data.expiration = time.Now().Add(time.Second * time.Duration(l.evictionTime))
		curNode.prev = nil
		l.head = curNode
		return l.head.data.value, nil
	}
	return "", errors.New("Key not found")
}

func (l *LRU) Display() {
	cur := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%v - %v\n", cur.data.key, cur.data.value)
		cur = cur.next
	}
	fmt.Println()
}

func (l *LRU) Evict() {
	if l.length == 0 {
		return
	}

	if time.Now().After(l.tail.data.expiration) {
		fmt.Printf("Removing key %v\n", l.tail.data.key)
		if l.length == 1 {
			l.head = nil
			l.tail = nil
			l.length = 0
			return
		}
		prev := l.tail.prev
		l.tail.prev = nil
		prev.next = nil
		l.tail = prev
		l.length--
	}
}
