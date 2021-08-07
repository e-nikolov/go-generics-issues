package main

func doFn(s string) {
	return
}

func Do[T any](do func(T)) {
	x := (interface{})(*new(T))
	do(x.(T))
}

func main() {
	Do(doFn)
}
