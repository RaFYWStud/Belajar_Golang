package contract

import "golang-tutorial/dto"

type Service struct {
	Intro IntroService
	ToDo ToDoService
}

// type exampleService interface {
// Code here
// }

type IntroService interface {
	GetIntro(introID int) (*dto.IntroResponse, error)
	CreateIntro(payload *dto.IntroRequest) (*dto.IntroResponse, error)
	UpdateIntro(id int, payload *dto.IntroRequest) (*dto.IntroResponse, error)
	DeleteIntro(id int) (*dto.IntroResponse, error)
}

type ToDoService interface {
    CreateToDo(payload *dto.ToDoRequest) (*dto.ToDoResponse, error)
    GetToDos() ([]dto.ToDoResponse, error)
	GetToDoByID(id int) (*dto.ToDoResponse, error)
	ReplaceToDo(id int, payload *dto.ToDoRequest) (*dto.ToDoResponse, error)
	UpdateToDo(id int, payload *dto.ToDoRequest) (*dto.ToDoResponse, error)
	DeleteToDo(id int) error
}
