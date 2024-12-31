package main

var ( //declarando tipos
	a bool    = true
	b int     = 10
	c string  = "teste"
	d float64 = 10.10
)

func x() {
	e := 12
	println(e)
}

func main() {
	println(a, b, c, d)
	x() //chamando uma função simples
}
