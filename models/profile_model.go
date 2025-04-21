package models

import (
	"gorm.io/gorm"
)

// gorm.Model definition
type Users struct {
	gorm.Model
	EmployeeID string    `json:"employee_id"`
	Name       string    `json:"name"`
	Lastname   string    `json:"lastname"`
	Birthday   string 	 `json:"birthday"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
}

type UsersRes struct {
	Name  string `json:"name"`
	EmployeeID string    `json:"employee_id"`
	Gen  string `json:"gen"`
}

type ResultData struct {
	Data  []UsersRes `json:"data"`
	Name  string    `json:"name"`
	Count int       `json:"count"`
	GenZ int		`json:"gen_z"`
	GenY int	`json:"gen_y"`
	GenX int	`json:"gen_x"`
	BabyBoomer int	`json:"baby_boomer"`
	GI_Generation int	`json:"gi_generation"`
}