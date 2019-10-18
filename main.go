package main

import (
	"fmt"

	"github.com/skmcgrail/go-mod-test-pkg1/services/serviceA"
	"mcgrail.xyz/v2/util"
)

func main() {
	err := serviceA.Operation()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Println(util.CurrentTime())
}
