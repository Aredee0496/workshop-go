package controllers

import (
	// "fmt"
	"go-workshop/database"
	m "go-workshop/models"
	// "log"
	// "regexp"
	// "strconv"
	"strings"

	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// users crud
func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users) //delete = null
	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var user []m.Users

	result := db.Find(&user, "employee_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&user)
}

func FindUser(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var user []m.Users

	result := db.Where(
		"employee_id = ? OR name = ? OR lastname = ?",
		search, search, search,
	).Find(&user)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&user)
}

func AddUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&user)
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user m.Users

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetUsersGen(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)

	genZ := 0
	genY := 0
	genX := 0
	BabyBoomer := 0
	GI_Generation := 0

	var dataResults []m.UsersRes
	for _, a := range users {
		Gen := ""
		if a.Age < 24 {
			Gen = "Gen Z"
			genZ = genZ + 1
		} else if a.Age >= 24 && a.Age <= 41 {
			Gen = "Gen Y"
			genY = genY + 1
		} else if a.Age >= 42 && a.Age <= 56 {
			Gen = "Gen X"
			genX = genX + 1
		} else if a.Age >= 57 && a.Age <= 75 {
			Gen = "Baby Boomer"
			BabyBoomer = BabyBoomer + 1
		} else if a.Age > 75 {
			Gen = "G.I. Generation"
			GI_Generation = GI_Generation + 1
		} else {
			Gen = "no gen"
		}

		d := m.UsersRes{
			Name:       a.Name,
			EmployeeID: a.EmployeeID,
			Gen:        Gen,
		}
		dataResults = append(dataResults, d)
	}
	r := m.ResultData{
		Data:          dataResults,
		Name:          "golang-workshop",
		Count:         len(users),
		GenZ:          genZ,
		GenY:          genY,
		GenX:          genX,
		BabyBoomer:    BabyBoomer,
		GI_Generation: GI_Generation,
	}
	return c.Status(200).JSON(r)
}
