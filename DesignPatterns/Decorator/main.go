package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	DoSomethingTimeConsuming(NewCtx())
}

func NewCtx() context.Context {
	dec := WithNocancelComponent(
		WithHashComponent(nil, "some_secret"),
	)

	return dec.Ctx(context.Background())
}

func DoSomethingTimeConsuming(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	fmt.Printf("%+v\n", ctx)
	time.Sleep(2 * time.Second)
}
