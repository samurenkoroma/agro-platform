package inmemory

import "sync"

type Registry struct {
	mu sync.RWMutex

	handlers map[string]any
}

func NewRegistry() *Registry {
	return &Registry{handlers: make(map[string]any)}
}

func (r *Registry) Register(name string, handler any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[name] = handler
}

func (r *Registry) Get(name string) (any, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	h, ok := r.handlers[name]

	return h, ok
}
