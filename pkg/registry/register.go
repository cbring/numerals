package registry

import (
	"sync"

	"github.com/gnirb/numerals/pkg/numerals"
	"github.com/gnirb/numerals/pkg/numerals/egyptian"
	"github.com/gnirb/numerals/pkg/numerals/etruscan"
	"github.com/gnirb/numerals/pkg/numerals/roman"
)

var (
	registryOnce     sync.Once
	registryInstance *registry
)

type registry struct {
	lock       sync.RWMutex
	converters map[string]numerals.Converter
}

func init() {
	Registry().Register("roman", roman.NewRomanConverter())
	Registry().Register("egyptian", egyptian.NewEgyptianConverter())
	Registry().Register("etruscan", etruscan.NewEtruscanConverter())
}

func Registry() *registry {
	registryOnce.Do(func() {
		registryInstance = &registry{
			converters: map[string]numerals.Converter{},
		}
	})

	return registryInstance
}

func (cr *registry) Register(name string, converter numerals.Converter) {
	cr.lock.Lock()
	defer cr.lock.Unlock()

	cr.converters[name] = converter
}

func (cr *registry) List() []string {
	cr.lock.RLock()
	defer cr.lock.RUnlock()

	var names []string
	for name := range cr.converters {
		names = append(names, name)
	}

	return names
}

func (cr *registry) Get(name string) numerals.Converter {
	cr.lock.RLock()
	defer cr.lock.RUnlock()

	return cr.converters[name]
}
