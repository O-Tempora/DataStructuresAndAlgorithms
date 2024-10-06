package main

import (
	"context"
)

// CtxCreatorComponent contains common behaviour of objects
// that create and decorate context.
type CtxCreatorComponent interface {
	Ctx(context.Context) context.Context
}
