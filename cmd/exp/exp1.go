package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, 123, "blue")
	value := ctx.Value(123)
	fmt.Println(value)
	fmt.Println(ctx)
}
