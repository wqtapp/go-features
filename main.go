package main

func main() {
	b := B{}
	b.what()
}

type A interface {
	what()
}

type B struct {
	A
}
