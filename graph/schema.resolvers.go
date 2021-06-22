package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "github.com/RobertMaulana/graphql-go/graph/generated"
	model1 "github.com/RobertMaulana/graphql-go/graph/model"
	"strconv"

	"github.com/RobertMaulana/graphql-go/graph/model"
	"github.com/RobertMaulana/graphql-go/internal/auth"
	"github.com/RobertMaulana/graphql-go/internal/pkg/db/model/links"
	"github.com/RobertMaulana/graphql-go/internal/pkg/db/model/users"
	"github.com/RobertMaulana/graphql-go/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model1.NewLink) (*model1.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	link.User = user
	linkId := link.Save()
	graphqlUser := &model.User{
		ID:       user.ID,
		Username: user.Username,
	}
	return &model.Link{ID: strconv.FormatInt(linkId, 10), Title: link.Title, Address: link.Address, User: graphqlUser}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model1.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model1.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model1.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model1.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return []*model.Link{}, fmt.Errorf("access denied")
	}
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		grahpqlUser := &model.User{
			ID:       link.User.ID,
			Username: link.User.Username,
		}
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address, User: grahpqlUser})
	}
	return resultLinks, nil
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
