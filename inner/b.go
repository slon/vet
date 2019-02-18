package main

import "github.com/slon/vet"

func main() {
	g := vet.G()
	g.Write([]byte("foo"))
}
