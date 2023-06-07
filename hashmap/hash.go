package hashmap

type hashMap[K comparable, V any] struct {
	hash *map[K]V
}

func New[K comparable, V any]() hashMap[K, V] {
	return hashMap[K, V]{
		hash: &map[K]V{},
	}
}

func (h hashMap[K, V]) Put(key K, value V) {
	(*h.hash)[key] = value
}

func (h hashMap[K, V]) Get(key K) V {
	return (*h.hash)[key]
}

func (h hashMap[K, V]) Contains(key K) bool {
	_, ok := (*h.hash)[key]
	return ok
}

func (h hashMap[K, V]) Size() int {
	return len(*h.hash)
}

func (h hashMap[K, V]) IsEmpty() bool {
	return len(*h.hash) == 0
}
