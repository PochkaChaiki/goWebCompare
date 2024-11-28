package todo

import (
	"fmt"
)

// type TodoList interface {
// 	GetList() []Todo
// 	GetTodo(id int) (Todo, error)
// 	CreateTodo(todo Todo) error
// 	UpdateTodo(todo Todo) error
// 	DeleteTodo(id int) error
// }

type TodoList struct {
	list   map[int]Todo
	lastId int
}

func New() *TodoList {
	return &TodoList{list: make(map[int]Todo), lastId: 1}
}

func (tl *TodoList) GetList() []Todo {
	list := []Todo{}
	for _, todo := range tl.list {
		list = append(list, todo)
	}
	return list
}

func (tl *TodoList) GetTodo(id int) (Todo, error) {
	if todo, ok := tl.list[id]; !ok {
		return Todo{}, fmt.Errorf("no todo in todolist with id=%d", id)
	} else {
		return todo, nil
	}
}

func (tl *TodoList) CreateTodo(todo Todo) error {
	todo.Id = tl.lastId
	tl.list[tl.lastId] = todo
	tl.lastId++
	return nil
}

func (tl *TodoList) UpdateTodo(todo Todo) error {
	if _, ok := tl.list[todo.Id]; !ok {
		return fmt.Errorf("no todo in todolist with id=%d use /api/create_todo", todo.Id)
	}
	tl.list[todo.Id] = todo
	return nil
}

func (tl *TodoList) DeleteTodo(id int) error {
	if _, ok := tl.list[id]; !ok {
		return fmt.Errorf("no todo in todolist with id=%d", id)
	}
	delete(tl.list, id)
	return nil
}
