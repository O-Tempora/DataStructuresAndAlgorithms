package main

import (
	"context"
	"crypto/sha256"
)

type Hash string

type HashComponent struct {
	c      CtxCreatorComponent
	secret Hash
}

func WithHashComponent(c CtxCreatorComponent, s string) HashComponent {
	return HashComponent{
		c:      c,
		secret: Hash(s),
	}
}

func (d HashComponent) Ctx(ctx context.Context) context.Context {
	hashedStr := sha256.Sum256([]byte(d.secret))
	ctx = context.WithValue(ctx, new(Hash), string(hashedStr[:32]))

	if d.c == nil {
		return ctx
	}

	return d.c.Ctx(ctx)
}
