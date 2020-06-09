package go_hello_module

import (
	"fmt"
	"github.com/morish000/go_hello_module/v2/hello"
	"github.com/morish000/go_hello_module/v2/world"
)

func HelloWorld() string {
	return fmt.Sprintf("%v %v %v", hello.Hello(), world.World(), "v2.1")
}
