package hashmap

import "testing"

func TestPutAndGet(t *testing.T) {
	h := New[string, string]()

	expected := "test"

	h.Put(expected, expected)

	actual := h.Get(expected)

	if expected != actual {
		t.Fatal("Expected", expected, "but got", actual)
	}
}

func TestSize(t *testing.T) {
	h := New[string, string]()

	key := "test"
	h.Put(key, key)

	expectedSize := 1
	actualSize := h.Size()

	if expectedSize != actualSize {
		t.Fatal("Expected", expectedSize, "but got", actualSize)
	}
}
