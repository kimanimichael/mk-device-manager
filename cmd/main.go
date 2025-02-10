package main

import (
	"fmt"
	"github.com/kimanimichael/mk-device-manager/internal/app"
)

func main() {
	fmt.Println("Welcome to MK Device Manager V1")

	app.NewServer()
}
