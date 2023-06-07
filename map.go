package standardmap

type Map[K comparable, V any] interface {
	Put(key K, value V)

	Get(key K) V

	Size() int

	Contains(key K) bool

	IsEmpty() bool
}
