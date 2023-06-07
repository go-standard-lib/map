package hashmap

type HashMap[K comparable, V any] struct {
	hash *map[K]V
}

func New[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{
		hash: &map[K]V{},
	}
}

func (h HashMap[K, V]) Put(key K, value V) {
	(*h.hash)[key] = value
}

func (h HashMap[K, V]) Get(key K) V {
	return (*h.hash)[key]
}

func (h HashMap[K, V]) Contains(key K) bool {
	_, ok := (*h.hash)[key]
	return ok
}

func (h HashMap[K, V]) Size() int {
	return len(*h.hash)
}

func (h HashMap[K, V]) IsEmpty() bool {
	return len(*h.hash) == 0
}

func (h HashMap[K, V]) Remove(key K) {
	delete(*h.hash, key)
}

func (h HashMap[K, V]) ToMap() map[K]V {
	return *h.hash
}

func (h HashMap[K, V]) Keys() []K {
	array := []K{}
	for k, _ := range *h.hash {
		array = append(array, k)
	}
	return array
}

func (h HashMap[K, V]) Values() []V {
	array := []V{}
	for _, v := range *h.hash {
		array = append(array, v)
	}
	return array
}
