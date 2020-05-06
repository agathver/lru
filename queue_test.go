package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Offer(t *testing.T) {
	t.Run("Should add an empty queue", func(t *testing.T) {
		q := NewQueue()
		q.Offer("hello")
		assert.Equal(t, 1, q.Size())

		assert.Equal(t, "hello", q.tail.value)
		assert.Equal(t, "hello", q.head.value)
	})

	t.Run("Should add to tail of a non-empty queue", func(t *testing.T) {
		q := NewQueue()
		q.Offer("Hello")
		q.Offer("World")

		assert.Equal(t, 2, q.Size())

		assert.Equal(t, "World", q.tail.value)
		assert.Equal(t, "Hello", q.head.value)
	})
}

func TestQueue_Poll(t *testing.T) {
	t.Run("Should poll from a Queue and remove the item", func(t *testing.T) {
		q := NewQueue()
		q.Offer("hello")
		q.Offer("World")

		assert.Equal(t, "hello", q.Poll())

		assert.Equal(t, 1, q.Size())
		assert.Equal(t,"World", q.head.value)
	})

	t.Run("Should poll from a Queue with single element and remove the item", func(t *testing.T) {
		q := NewQueue()
		q.Offer("hello")

		assert.Equal(t, "hello", q.Poll())

		assert.Equal(t, 0, q.Size())
		assert.Nil(t, q.head)
		assert.Nil(t, q.tail)
	})

	t.Run("Should return nil when queue is empty", func(t *testing.T) {
		q := NewQueue()
		q.Offer("hello")

		assert.Equal(t, "hello", q.Poll())

		assert.Nil(t, q.Poll())
	})
}
