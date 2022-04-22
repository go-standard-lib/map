package map

type Map[K comparable, V any] interface {
	Put(key K, value V)

	Get(key K) V

	Size() int

	IsEmpty() bool
}