package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	parserservice "github.com/fracartdev/testing/app/parser"
	"github.com/fracartdev/testing/graph/generated"
	"github.com/fracartdev/testing/graph/model"
)

func (r *queryResolver) Order(ctx context.Context, id string) (*model.Order, error) {
	var order *model.Order
	var report *[]model.Order
	encJSON := parserservice.Parse()
	err := json.Unmarshal(encJSON, &report)
	fmt.Println(report)
	return order, err
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
