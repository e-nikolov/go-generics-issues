package pkg

type Doer[T any] interface {
	Do() T
}
