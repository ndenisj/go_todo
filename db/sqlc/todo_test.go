package db

import (
	"context"
	"testing"
	"time"

	"github.com/ndenisj/go_todo/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTodo(t *testing.T) Todo {
	arg := CreateTodoParams{
		Owner:   utils.RandomOwner(),
		Title:   utils.RandomTitle(),
		Content: utils.RandomContent(),
	}

	todo, err := testQueries.CreateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, todo.Owner, arg.Owner)
	require.Equal(t, todo.Title, arg.Title)
	require.Equal(t, todo.Content, arg.Content)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreatedAt)

	return todo
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
}

func TestGetTodo(t *testing.T) {
	// create todo
	todo1 := createRandomTodo(t)
	// get the todo
	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.Title, todo2.Title)
	require.Equal(t, todo1.Content, todo2.Content)

	require.WithinDuration(t, todo1.CreatedAt, todo2.CreatedAt, time.Second)
}

func TestListTodo(t *testing.T) {
	for i := 0; i <= 10; i++ {

		createRandomTodo(t)
	}

	arg := ListTodosParams{
		Limit:  5,
		Offset: 5,
	}

	todos, err := testQueries.ListTodos(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, todos, 5)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
	}
}

func TestUpdateTodo(t *testing.T) {
	todo1 := createRandomTodo(t)

	arg := UpdateTodoParams{
		ID:      todo1.ID,
		Title:   utils.RandomTitle(),
		Content: utils.RandomContent(),
	}

	todo2, err := testQueries.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo2.ID, arg.ID)
	require.Equal(t, todo2.Title, arg.Title)
	require.Equal(t, todo2.Content, arg.Content)

	require.NotEqual(t, todo1.Title, todo2.Title)
	require.NotEqual(t, todo1.Content, todo2.Content)
}

func TestDeleteTodo(t *testing.T) {
	todo := createRandomTodo(t)

	err := testQueries.DeleteTodo(context.Background(), todo.ID)

	require.NoError(t, err)
}
