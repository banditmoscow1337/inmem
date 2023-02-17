package dadb

import "sync"

type Storage struct {
	mux sync.RWMutex
	db  map[string]interface{}
}

func New(dbs ...map[string]interface{}) *Storage {
	db := make(map[string]interface{})
	if len(dbs) == 1 {
		db = dbs[0]
	}
	store := &Storage{
		db: db,
	}
	return store
}

func (s *Storage) Get(key string) interface{} {
	s.mux.RLock()
	v, ok := s.db[key]
	s.mux.RUnlock()
	if !ok {
		return nil
	}
	return v
}

func (s *Storage) Set(key string, val interface{}) {
	s.mux.Lock()
	s.db[key] = val
	s.mux.Unlock()
}

func (s *Storage) Delete(key string) {
	s.mux.Lock()
	delete(s.db, key)
	s.mux.Unlock()
}

func (s *Storage) Reset() {
	ndb := make(map[string]interface{})
	s.mux.Lock()
	s.db = ndb
	s.mux.Unlock()
}

func (s *Storage) Conn() map[string]interface{} {
	return s.db
}
