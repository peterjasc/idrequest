package csv

import "sync"

type SynchronizedMap struct {
	sync.RWMutex
	unsynchronizedMap map[string][2]string
}

func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		unsynchronizedMap: make(map[string][2]string),
	}
}

func (sm *SynchronizedMap) Load(key string) (result [2]string, ok bool) {
	sm.RLock()
	result, ok = sm.unsynchronizedMap[key]
	sm.RUnlock()
	return result, ok
}

func (sm *SynchronizedMap) UpdateMap(newMap map[string][2]string) {
	sm.Lock()
	sm.unsynchronizedMap = newMap
	sm.Unlock()
}
