package main

import (
	"context"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/boot"
)

func main() {
	ctx := context.Background()
	err := boot.Boot(ctx)
	if err != nil {
		fmt.Printf("Boot failed: %v", err)
	}
}
