package gin_graphql

import (
	"context"
	"gin_graphql/internal/models"
	connection "gin_graphql/pkg/database"
)

var db = connection.GetInstance()

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*models.User, error) {
	user := models.User{
		Name: input.Name,
	}
	db.Create(&user)
	return &user, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*models.Todo, error) {
	todo := models.Todo{
		Text:   input.Text,
		UserID: input.UserID,
		Done:   false,
	}
	db.Create(&todo)
	return &todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input EditTodo) (*models.Todo, error) {
	todo := models.Todo{ID: input.ID}
	db.First(&todo)
	todo.Text = input.Text
	db.Model(&models.Todo{}).Update(&todo)

	return &todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int) (*models.Todo, error) {
	todo := models.Todo{
		ID: input,
	}
	db.First(&todo)
	db.Delete(&todo)
	return &todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, input *FetchTodo) (*models.Todo, error) {
	var todo models.Todo
	db.Preload("User").First(&todo, input.ID)
	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	var todos []models.Todo
	db.Preload("User").Find(&todos)
	return todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	var users []models.User
	db.Find(&users)
	return users, nil
}
