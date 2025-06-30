package main

import (
	"context"
	"fmt"
)

type ownContextKey string

const key ownContextKey = "key"

type someoneElsesContextKey string

const someoneElsesKey someoneElsesContextKey = "key"

func main() {
	ctx := context.Background()

	ctx2 := context.WithValue(ctx, key, "value")
	ctx3 := context.WithValue(ctx2, someoneElsesKey, "value2")
	fmt.Println("ctx2", ctx2)
	fmt.Println("ctx3", ctx3.Value(key))
	fmt.Println("ctx3", ctx3.Value(someoneElsesKey))
}

func SetKey(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, key, v)
}

func GetKey(ctx context.Context) (string, bool) {
	if v, ok := ctx.Value(key).(string); ok {
		return v, true
	}
	return "", false
}
