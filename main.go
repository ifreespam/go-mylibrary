package main

import (
	"fmt"
	"github.com/ifreespam/go-mylibrary/sdk"
)

func main() {
	client := sdk.NewClient("0.0.0.0", 1234)
	client.PrintInfo()

	fmt.Println("Initialised successfully!")
}
