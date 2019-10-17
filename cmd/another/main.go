package main

import (
	"fmt"

	"github.com/skmcgrail/go-mod-test-pkg1/services/serviceA/v2"
)

func main() {
	b, err := serviceA.Operation()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(b)
}
