// Fun with Context...

package main

import (
	stdctx "context"
	"fmt"

	"github.com/imattf/galere/context"
	"github.com/imattf/galere/models"
)

// type ctxKey string

// const favoriteColorKey ctxKey = "favorite-color"

func main() {
	// ctx := context.Background()
	ctx := stdctx.Background()

	user := models.User{
		Email: "bob@aol.com",
	}

	// our stored value in the context
	// ctx = context.WithValue(ctx, favoriteColorKey, 123)
	ctx = context.WithUser(ctx, &user)

	// a different stored value in the same context, with same name-ish
	// ctx = context.WithValue(ctx, "favorite-color", "red")

	// value1 := ctx.Value(favoriteColorKey)
	// value2 := ctx.Value("favorite-color")
	// anyValue := ctx.Value(favoriteColorKey)
	retrievedUser := context.User(ctx)

	// fmt.Println(value1)
	// fmt.Println(value2)

	// This .(string) is type assertion. Use ok to avoid run-time panic
	// stringValue, ok := anyValue.(string)
	// if !ok {
	// 	fmt.Println(anyValue, "is not a string")
	// 	return
	// }

	// // fmt.Println(ctx)

	// fmt.Println(stringValue, "is a string")
	fmt.Println(retrievedUser.Email)

}
