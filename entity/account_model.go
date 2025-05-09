package entity

type Account struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Username  string `gorm:"column:username;type:varchar(255);not null"`
	Email     string `gorm:"column:email;type:varchar(255);not null;unique"`
	Password  string `gorm:"column:password;type:varchar(255);not null"`
	CreatedAt string `gorm:"column:created_at;type:timestamp;not null;default:now()"`
	UpdatedAt string `gorm:"column:updated_at;type:timestamp;not null;default:now()"`
}

func (e *Account) TableName() string {
	return "account"
}