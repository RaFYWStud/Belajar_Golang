package repository

import (
    "fmt"
    "golang-tutorial/contract"
    "golang-tutorial/entity"

    "gorm.io/gorm"
)

type todoRepo struct {
    db *gorm.DB
}

func implToDoRepository(db *gorm.DB) contract.ToDoRepository {
    return &todoRepo{
        db: db,
    }
}

func (r *todoRepo) GetToDos() ([]entity.ToDo, error) {
    var todos []entity.ToDo
    err := r.db.Table("to_do_list").Find(&todos).Error
    if err != nil {
        return nil, err
    }
    return todos, nil
}

func (r *todoRepo) GetToDoByID(id int) (*entity.ToDo, error) {
    var todo entity.ToDo
    err := r.db.Table("to_do_list").Where("id = ?", id).First(&todo).Error
    if err != nil {
        return nil, err
    }
    return &todo, nil
}

func (r *todoRepo) CreateToDo(todo *entity.ToDo) error {
	if todo.ID == 0 {
        return fmt.Errorf("ID is required")
    }
    return r.db.Table("to_do_list").Create(todo).Error
}

func (r *todoRepo) ReplaceToDo(id int, todo *entity.ToDo) error {
    return r.db.Table("to_do_list").Where("id = ?", id).Updates(todo).Error
}

func (r *todoRepo) UpdateToDo(id int, updates map[string]interface{}) error {
    return r.db.Table("to_do_list").Where("id = ?", id).Updates(updates).Error
}

func (r *todoRepo) DeleteToDo(id int) error {
    return r.db.Table("to_do_list").Where("id = ?", id).Delete(&entity.ToDo{}).Error
}

