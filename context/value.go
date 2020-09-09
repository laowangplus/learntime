package main

import (
	"context"
	"fmt"
)

func f(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "foo", -6)
	fmt.Println(ctx.Value("foo"))
	return ctx
}

func main() {
	ctx := context.Background()
	ctx = f(ctx)
	fmt.Println(ctx.Value("foo"))
}
