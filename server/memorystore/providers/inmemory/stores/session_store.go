package stores

import (
	"os"
	"sync"

	"github.com/authorizerdev/authorizer/server/constants"
)

// SessionStore struct to store the env variables
type SessionStore struct {
	mutex sync.Mutex
	store map[string]map[string]string
}

// NewSessionStore create a new session store
func NewSessionStore() *SessionStore {
	return &SessionStore{
		mutex: sync.Mutex{},
		store: make(map[string]map[string]string),
	}
}

// Get returns the value of the key in state store
func (s *SessionStore) Get(key, subKey string) string {
	return s.store[key][subKey]
}

// Set sets the value of the key in state store
func (s *SessionStore) Set(key string, subKey, value string) {
	if os.Getenv("ENV") != constants.TestEnv {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if _, ok := s.store[key]; !ok {
		s.store[key] = make(map[string]string)
	}
	s.store[key][subKey] = value
}

// RemoveAll all values for given key
func (s *SessionStore) RemoveAll(key string) {
	if os.Getenv("ENV") != constants.TestEnv {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	delete(s.store, key)
}

// Remove value for given key and subkey
func (s *SessionStore) Remove(key, subKey string) {
	if os.Getenv("ENV") != constants.TestEnv {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	if _, ok := s.store[key]; ok {
		delete(s.store[key], subKey)
	}
}

// Get all the values for given key
func (s *SessionStore) GetAll(key string) map[string]string {
	if _, ok := s.store[key]; !ok {
		s.store[key] = make(map[string]string)
	}
	return s.store[key]
}