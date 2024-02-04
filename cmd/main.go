package main

import (
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/boot"
)

func main() {
	rt, err := boot.Boot()
	if err != nil {
		fmt.Printf("Boot failed: %v", err)
	}

	rt.StartServe()
}
