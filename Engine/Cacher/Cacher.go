package Cacher

import (
	"sync"
)

// add metadata in the future ?? -> map[string]struct{...}
// make it hold different data types

type KeyValueStore struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

func (kvs *KeyValueStore) Get(key string) (string, error) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()
	entry := kvs.store[key]
	return entry, nil
}

// locks map and sets a key.
func (kvs *KeyValueStore) Set(key string, value string) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	kvs.store[key] = value
	return nil
}

// same as Set but will NOT overwrite if a key exists
func (kvs *KeyValueStore) SafeSet(key string, value string) error {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	_, exists := kvs.store[key]
	if !exists {
		kvs.store[key] = value
	}
	return nil
}

// check if key exists
func (kvs *KeyValueStore) Exists(key string) (bool, error) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()
	_, exists := kvs.store[key]
	return exists, nil
}

// deletes key
func (kvs *KeyValueStore) Delete(key string) (bool, error) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	delete(kvs.store, key)
	return true, nil
}
