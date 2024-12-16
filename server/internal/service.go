package internal

import "errors"

type MemoryStore struct {
	m map[string]string //todo: make map[string]interface{}
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		m: make(map[string]string),
	}
}

func (ms *MemoryStore) Put(key, val string) error {
	ms.m[key] = val
	return nil
}

func (ms *MemoryStore) Get(key string) (string, error) {
	if v, ok := ms.m[key]; ok {
		return v, nil
	}

	return "", errors.New("key doesn't exist")
}
