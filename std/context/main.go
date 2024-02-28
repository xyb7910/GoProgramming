package main

import (
	"context"
	"fmt"
)

func main() {
	// create a new context
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "jack")
	Get(ctx)
}

func Get(ctx context.Context) {
	fmt.Println(ctx.Value("name"))
}
