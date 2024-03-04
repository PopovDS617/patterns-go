package mutexes

import (
	"net"
	"sync"
)

var (
	service map[string]net.Addr
	m       sync.RWMutex
)

func PostService(name string, addr net.Addr) {
	m.Lock()
	defer m.Unlock()

	service[name] = addr
}
func GetService(name string) net.Addr {
	m.RLock()
	defer m.RUnlock()

	return service[name]
}
