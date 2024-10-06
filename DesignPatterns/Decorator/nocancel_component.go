package main

import (
	"context"
)

type NocancelComponent struct {
	c CtxCreatorComponent
}

func WithNocancelComponent(c CtxCreatorComponent) NocancelComponent {
	return NocancelComponent{
		c: c,
	}
}

func (d NocancelComponent) Ctx(ctx context.Context) context.Context {
	ctx = context.WithoutCancel(ctx)
	if d.c == nil {
		return ctx
	}
	return d.c.Ctx(ctx)
}
