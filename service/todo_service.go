package service

import (
    "golang-tutorial/contract"
    "golang-tutorial/dto"
    "golang-tutorial/entity"
)

type ToDoService struct {
    todoRepository contract.ToDoRepository
}

func implToDoService(repo *contract.Repository) contract.ToDoService {
    return &ToDoService{
        todoRepository: repo.ToDo,
    }
}

func (s *ToDoService) CreateToDo(payload *dto.ToDoRequest) (*dto.ToDoResponse, error) {
    todo := &entity.ToDo{
        ID:   payload.ID,
        Nama: payload.Nama,
        Hari: payload.Hari,
        ToDo: payload.ToDo,
    }

    err := s.todoRepository.CreateToDo(todo)
    if err != nil {
        return nil, err
    }

    return &dto.ToDoResponse{
        StatusCode: 201,
        Message:    "Berhasil membuat To-Do",
        Data: dto.ToDoData{
            ID:   todo.ID,
            Nama: todo.Nama,
            Hari: todo.Hari,
            ToDo: todo.ToDo,
        },
    }, nil
}

func (s *ToDoService) GetToDos() ([]dto.ToDoResponse, error) {
    todos, err := s.todoRepository.GetToDos()
    if err != nil {
        return nil, err
    }

    var responses []dto.ToDoResponse
    for _, todo := range todos {
        responses = append(responses, dto.ToDoResponse{
            StatusCode: 200,
            Message:    "Berhasil mendapatkan To-Do",
            Data: dto.ToDoData{
                ID:   todo.ID,
                Nama: todo.Nama,
                Hari: todo.Hari,
                ToDo: todo.ToDo,
            },
        })
    }

    return responses, nil
}

func (s *ToDoService) GetToDoByID(id int) (*dto.ToDoResponse, error) {
    todo, err := s.todoRepository.GetToDoByID(id)
    if err != nil {
        return nil, err
    }

    return &dto.ToDoResponse{
        StatusCode: 200,
        Message:    "Berhasil mendapatkan To-Do",
        Data: dto.ToDoData{
            ID:   todo.ID,
            Nama: todo.Nama,
            Hari: todo.Hari,
            ToDo: todo.ToDo,
        },
    }, nil
}