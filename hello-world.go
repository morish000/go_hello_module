package go_hello_module

import (
	"fmt"
	"github.com/morish000/go_hello_module/hello"
	"github.com/morish000/go_hello_module/world"
)

func HelloWorld() string {
	return fmt.Sprintf("%v %v", hello.Hello(), world.World())
}