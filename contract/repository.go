package contract

import "golang-tutorial/entity"

type Repository struct {
	Intro IntroRepository
	ToDo ToDoRepository
	Account AccountRepository
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
	ReplaceToDo(id int, todo *entity.ToDo) error
	UpdateToDo(id int, updates map[string]interface{}) error
	DeleteToDo(id int) error
}

type AccountRepository interface {
	GetAccount(id int) (*entity.Account, error)
	CreateAccount(account *entity.Account) error
	CheckEmail(email string) (bool, error)
	GetAccountByEmail(email string) (*entity.Account, error)
	GetAccountByUsername(username string) (*entity.Account, error)
}