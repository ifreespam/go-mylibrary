package main

import "github.com/ifreespam/go-mylibrary/sdk"

func main() {
	client := sdk.NewClient("0.0.0.0", 1234)
	client.PrintInfo()
}
