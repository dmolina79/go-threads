package lstoreds

// cache abstracts all methods we access from ARCCache, to enable alternate
// implementations such as a no-op one.
type cache interface {
	Get(key interface{}) (value interface{}, ok bool)
	Add(key, value interface{})
	Remove(key interface{})
	Contains(key interface{}) bool
	Peek(key interface{}) (value interface{}, ok bool)
	Keys() []interface{}
}

// noopCache is a dummy implementation that's used when the cache is disabled.
type noopCache struct {
}

var _ cache = (*noopCache)(nil)

func (*noopCache) Get(key interface{}) (value interface{}, ok bool) {
	return nil, false
}

func (*noopCache) Add(key, value interface{}) {
}

func (*noopCache) Remove(key interface{}) {
}

func (*noopCache) Contains(key interface{}) bool {
	return false
}

func (*noopCache) Peek(key interface{}) (value interface{}, ok bool) {
	return nil, false
}

func (*noopCache) Keys() (keys []interface{}) {
	return keys
}
