package inner

import "github.com/slon/vet"

func bar() {
	g := vet.G()
	g.Write([]byte("foo"))
}
