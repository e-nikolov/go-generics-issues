package main

func doFn(s string) {
	return
}

func Do[T any](do func(string)) {
	_ = (interface{})(*new(string)).(string)
}

func main() {
	Do[string](doFn)
}
