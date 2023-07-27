package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	graph "github.com/minkezhang/truffle/api/graphql/generated"
	"github.com/minkezhang/truffle/api/graphql/generated/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Entry is the resolver for the Entry field.
func (r *mutationResolver) Entry(ctx context.Context, input *model.MutateEntryInput) (*model.Entry, error) {
	return NewEntry(input)
}

// Entry is the resolver for the Entry field.
func (r *queryResolver) Entry(ctx context.Context, input *model.QueryEntryInput) ([]*model.Entry, error) {
	return nil, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: fmt.Sprintf("Entry not found: %s", *input.ID),
		Extensions: map[string]interface{}{
			"code": 404,
		},
	}
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
