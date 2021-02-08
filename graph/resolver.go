package graph

import (
	mycrud "github.com/fracartdev/testing/app/internal"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MyCRUD mycrud.MyCRUD
}
