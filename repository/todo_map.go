package repository

import (
	"errors"
	"sync"
	"todo/domain"
)

// インターフェースを返す
type todoRepository struct {
	m sync.Map
}

func NewSyncMapTodoRepository() domain.TodoRepository {
	return &todoRepository{}
}

func (t *todoRepository) AllGet() ([]domain.Todo, error) {
	var todos []domain.Todo

	t.m.Range(func(key interface{}, value interface{}) bool {
		todos = append(
			todos,
			value.(domain.Todo),
		)
		return true
	})
	return todos, nil
}

func (t *todoRepository) StatusUpdate(id int) error {
	r, ok := t.m.LoadAndDelete(id)
	if !ok {
		return errors.New("data is not exist")
	}
	newTodo := r.(domain.Todo)
	if newTodo.Completed {
		newTodo.Completed = false
	} else {
		newTodo.Completed = true
	}

	t.Store(newTodo)
	return nil
}

func (t *todoRepository) Store(todo domain.Todo) error {
	t.m.Store(todo.ID, todo)
	return nil
}

func (t *todoRepository) Delete(id int) error {
	t.m.Delete(id)
	return nil
}
