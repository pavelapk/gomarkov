package gomarkov

import (
	"runtime"
	"sync"
)

type spool struct {
	stringMap map[string]int
	intMap    []string
	sync.RWMutex
}

func (s *spool) add(str string) int {
	s.RLock()
	index, ok := s.stringMap[str]
	s.RUnlock()
	if ok {
		return index
	}
	s.Lock()
	defer s.Unlock()
	index, ok = s.stringMap[str]
	if ok {
		return index
	}
	s.intMap = append(s.intMap, str)
	index = len(s.intMap) - 1
	s.stringMap[str] = index
	return index
}

func (s *spool) get(str string) (int, bool) {
	s.RLock()
	index, ok := s.stringMap[str]
	s.RUnlock()
	return index, ok
}

func (s *spool) getClosest(str string) (int, bool) {
	s.RLock()
	index := ConcurrentSearchClosest(s.intMap, str, runtime.NumCPU())
	s.RUnlock()
	ok := index >= 0
	return index, ok
}
