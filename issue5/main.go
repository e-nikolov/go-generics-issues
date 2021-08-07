package main

import "github.com/e-nikolov/go-generics-issues/issue5/pkg"

func Do[T any](doer pkg.Doer[T]) {
	doer.Do()
}

func main() {
}

type Doer[T any] interface {
	Do() T
}
