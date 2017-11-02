package stringset

import (
	"sync"
)

type StringSet struct {
	data map[string]bool
	sync.RWMutex
}

func New() *StringSet {
	return &StringSet{data: make(map[string]bool)}
}

func NewFromSlice(slc []string) *StringSet {
	newSet := New()
	for _, val := range slc {
		newSet.Add(val)
	}

	return newSet
}

func (s *StringSet) Add(val string) {
	s.Lock()
	defer s.Unlock()
	s.data[val] = true
}

func (s *StringSet) Remove(val string) {
	s.Lock()
	defer s.Unlock()
	if s.data[val] {
		delete(s.data, val)
	}
}

func (s *StringSet) Slice() []string {
	s.RLock()
	defer s.RUnlock()
	newSlice := make([]string, len(s.data))

	var i = 0
	for key := range s.data {
		newSlice[i] = key
		i++
	}

	return newSlice
}

func (s *StringSet) Union(sSet *StringSet) *StringSet {
	s.RLock()
	sSet.RLock()
	defer sSet.RUnlock()
	defer s.RUnlock()

	newSet := StringSet{data: s.data}
	for val := range sSet.data {
		newSet.data[val] = true
	}

	return &newSet
}

func (s *StringSet) Intersection(sSet *StringSet) *StringSet {
	s.RLock()
	sSet.RLock()
	defer s.RUnlock()
	defer sSet.RUnlock()

	newSet := New()
	for val := range s.data {
		if sSet.data[val] {
			newSet.data[val] = true
		}
	}

	return newSet
}

func (s *StringSet) Difference(sSet *StringSet) *StringSet {
	s.RLock()
	sSet.RLock()
	defer s.RUnlock()
	defer sSet.RUnlock()

	newSet := New()

	for val := range s.data {
		if !sSet.data[val] {
			newSet.data[val] = true
		}
	}

	return newSet
}
