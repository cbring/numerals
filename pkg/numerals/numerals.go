package numerals

import "sync"

type Numeral rune
type Number string

type Converter interface {
	Format(int) Number
	Parse(Number) int
}

var (
	registryOnce     sync.Once
	registryInstance *registry
)

type registry struct {
	lock       sync.RWMutex
	converters map[string]Converter
}

func Registry() *registry {
	registryOnce.Do(func() {
		registryInstance = &registry{
			converters: map[string]Converter{},
		}
	})

	return registryInstance
}

func (cr *registry) Register(name string, converter Converter) {
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

func (cr *registry) Get(name string) Converter {
	cr.lock.RLock()
	defer cr.lock.RUnlock()

	return cr.converters[name]
}
