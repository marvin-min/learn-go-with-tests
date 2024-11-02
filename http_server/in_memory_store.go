package main

import "sync"

type InMemoryStore struct {
	store map[string]int
	lock  sync.RWMutex
}

func (i *InMemoryStore) GetLeague() League {
	var league League
	for k, v := range i.store {
		league = append(league, Player{k, v})
	}
	return league
}
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{map[string]int{}, sync.RWMutex{}}
}
func (i *InMemoryStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++

}

func (i *InMemoryStore) GetPlayerScore(name string) int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.store[name]
}
