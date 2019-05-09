package gin_graphql

import (
	"context"
	"gin_graphql/internal/models"
	connection "gin_graphql/pkg/database"
)

var db = connection.GetInstance()

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []models.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	todo := models.Todo{
		Text:   input.Text,
		UserID: input.UserID,
	}

	db.Model(&models.ExtendedTodo{}).Create(&todo)
	return &todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	return r.todos, nil
}
