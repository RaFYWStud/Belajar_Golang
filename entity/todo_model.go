package entity

type ToDo struct {
    ID   int    `gorm:"column:id;primaryKey;not null"`
    Nama string `gorm:"column:nama;type:varchar(255);not null"`
    Hari string `gorm:"column:hari;type:varchar(50);not null"`
    ToDo string `gorm:"column:todo;type:text;not null"`
}

func (ToDo) TableName() string {
    return "to_do_list"
}