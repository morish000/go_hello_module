package hello

import (
	"fmt"
	"go-hello-module/hello"
	"go-hello-module/world"
)

func HelloWorld() string {
	return fmt.Sprintf("%v %v", hello.Hello(), world.World())
}