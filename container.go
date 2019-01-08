package container

import "sync"

// Container is the foundation of the container store
type Container struct {
	rwmutex  sync.RWMutex
	services map[string]interface{}
}

// NewContainer is factory method for the container
func NewContainer() *Container {
	var rwmutex sync.RWMutex
	services := make(map[string]interface{})
	return &Container{
		rwmutex,
		services,
	}
}

// Get returns a service by name
func (c *Container) Get(name string) (object interface{}, ok bool) {
	c.rwmutex.RLock()
	object, ok = c.services[name]
	c.rwmutex.RUnlock()
	return object, ok
}

// Add adds object to the internal store
func (c *Container) Add(name string, object interface{}) {
	c.rwmutex.Lock()
	c.services[name] = object
	c.rwmutex.Unlock()
}

// Remove service from internal store
func (c *Container) Remove(name string) {
	c.rwmutex.Lock()
	delete(c.services, name)
	c.rwmutex.Unlock()
}
