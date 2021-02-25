package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Storage struct {
	source string
	body   map[string]interface{}
	lock   sync.RWMutex
}

func (s *Storage) Len() int {
	return len(s.body)
}

func (s *Storage) write() {
	s.lock.Lock()
	defer s.lock.Unlock()

	writeStorage(s.source, s.body)
}

func (s *Storage) Get(key string) interface{} {
	return s.body[key]
}

func (s *Storage) GetAll() interface{} {
	var data sync.Map
	data.Range(func(key, value interface{}) bool {
		s.body[fmt.Sprint(key)] = value
		return true
	})
	b, err := json.MarshalIndent(s.body, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func (s *Storage) Set(key string, val interface{}) {
	s.body[key] = val
	s.write()
}

func (s *Storage) Del(key string) {
	delete(s.body, key)
	s.write()
}

func (s *Storage) Exist(key string) bool {
	_, ok := s.body[key]
	return ok
}

func (s *Storage) Rename(oldkey string, newkey string) {
	val := s.Get(oldkey)
	s.Set(newkey, val)
	s.Del(oldkey)
	s.write()
}

func (s *Storage) Clear() {
	s.body = make(map[string]interface{})
	s.write()
}

func New(source string) *Storage {
	content, err := ioutil.ReadFile(source)
	if err != nil {
		if os.IsNotExist(err) {
			content = createStorage(source)
		} else {
			panic(err)
		}
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(content, &parsed); err != nil {
		panic(err)
	}

	return &Storage{
		source: source,
		body:   parsed,
	}
}

func createStorage(source string) []byte {
	f, err := os.Create(source)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := []byte("{}\n")
	f.Write(content)
	return content
}

func writeStorage(source string, content map[string]interface{}) {
	result, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(source, result, 0644); err != nil {
		panic(err)
	}
}
