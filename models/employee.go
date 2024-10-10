package models

type Employee struct {
    ID       uint	`gorm:"primaryKey;autoIncrement" json:"id"`
    Name     string `gorm:"type:varchar(255)" validate:"required" json:"name"`
    Position string `gorm:"type:varchar(255)" validate:"required" json:"position"`
    Salary   int    `gorm:"type:int" validate:"required,gt=0" json:"salary"`
}
