package main

import (
	"fmt"

	"github.com/skmcgrail/go-mod-test-pkg1/services/serviceA"
)

func main() {
	err := serviceA.Operation()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
