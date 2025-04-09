package contract

import "golang-tutorial/entity"

type Repository struct {
	Intro IntroRepository
	ToDo ToDoRepository
}

// type exampleRepository interface {
// Code here
// }

type IntroRepository interface {
	CreateIntro(intro *entity.Intro) error
	GetIntro(id int) (*entity.Intro, error)
	UpdateIntro(id int, intro *entity.Intro) error
	DeleteIntro(id int) error
}

type ToDoRepository interface {
    GetToDos() ([]entity.ToDo, error)
    CreateToDo(todo *entity.ToDo) error
	GetToDoByID(id int) (*entity.ToDo, error)
}