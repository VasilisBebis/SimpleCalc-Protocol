package main

import (
	"fmt"
	"github.com/VasilisBebis/SimpleCalc-Protocol/internal/client"
	"github.com/VasilisBebis/SimpleCalc-Protocol/internal/server"
)

func main() {
	fmt.Println("Hello World")
	client.HelloClient()
	server.HelloServer()
}
