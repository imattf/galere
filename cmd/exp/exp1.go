package main

import (
	"context"
	"fmt"
)

type ctxKey string

const favoriteColorKey ctxKey = "favorite-color"

func main() {
	ctx := context.Background()

	// our stored value in the context
	ctx = context.WithValue(ctx, favoriteColorKey, 123)

	// a different stored value in the same context, with same name-ish
	// ctx = context.WithValue(ctx, "favorite-color", "red")

	// value1 := ctx.Value(favoriteColorKey)
	// value2 := ctx.Value("favorite-color")
	anyValue := ctx.Value(favoriteColorKey)

	// fmt.Println(value1)
	// fmt.Println(value2)

	// This .(string) is type assertion. Use ok to avoid run-time panic
	stringValue, ok := anyValue.(string)
	if !ok {
		fmt.Println(anyValue, "is not a string")
		return
	}

	// fmt.Println(ctx)

	fmt.Println(stringValue, "is a string")
}
