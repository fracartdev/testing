package mycrud

import "github.com/fracartdev/testing/graph/model"

type MyCRUD interface {
	ReadOrder(id string) (*model.Order, error)
}
