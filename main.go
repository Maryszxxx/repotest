package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Esse é o programa do exercício de compilação cruzada. foi um compilado num windows/amd64 e agora está rodando num sistema: ", runtime.GOARCH, runtime.GOOS)

}
