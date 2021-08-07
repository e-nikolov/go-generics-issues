package main

func Do[T any](do func() (T, T)) {
	_ = func() (T, T) {
		return do()
	}
}

func main() {
	Do[int](func() (int, int) {
		return 3, 3
	})
}
