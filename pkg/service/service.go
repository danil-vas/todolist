package service

import (
	"todolist"
	"todolist/pkg/repository"
)

type Authorization interface {
	CreateUser(user todolist.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todolist.TodoList) (int, error)
	GetAll(userId int) ([]todolist.TodoList, error)
	GetById(userId int, listId int) (todolist.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todolist.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item todolist.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todolist.TodoItem, error)
	GetById(userId, itemId int) (todolist.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todolist.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
