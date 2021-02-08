package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	api "github.com/fracartdev/testing/app/internal"
	parserservice "github.com/fracartdev/testing/app/parser"
	"github.com/fracartdev/testing/graph/generated"
	"github.com/fracartdev/testing/graph/model"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, order *model.InputOrder) (*model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	order, err := api.ReadOrder(id)

	return &model.Order{
		ID:         order.ID,
		Item:       order.Item,
		User:       order.User,
		Quantity:   order.Quantity,
		City:       order.City,
		Department: order.Department,
		Price:      order.Price,
	}, err
}

func (r *queryResolver) Report(ctx context.Context, id string) (*model.Report, error) {
	var collection []*model.Order
	encJSON := parserservice.Parse()

	err := json.Unmarshal(encJSON, &collection)

	return &model.Report{
		ID:        "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		OrderList: collection,
	}, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
