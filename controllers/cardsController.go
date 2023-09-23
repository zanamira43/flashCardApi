package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/flashcardApi/database"
	"github.com/zanamira43/flashcardApi/models"
)

// 1- get all decks
func GetALLCards(c *fiber.Ctx) error {
	var cards []models.Card
	database.DB.Find(&cards)
	return c.Status(200).JSON(fiber.Map{
		"cards": cards,
	})
}

// 2- get create new deck
func CreateNewCard(c *fiber.Ctx) error {
	var card models.Card

	if err := c.BodyParser(&card); err != nil {
		return err
	}

	database.DB.Create(&card)

	return c.Status(201).JSON(fiber.Map{
		"card": card,
	})

}

// 3- get single deck
func GetCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		panic(err)
	}
	var card models.Card
	database.DB.Where("id = ?", id).First(&card)

	return c.Status(200).JSON(fiber.Map{
		"card": card,
	})

}

// 4- update single deck
func UpdateCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		panic(err)
	}

	card := models.Card{
		ID: uint(id),
	}

	if err := c.BodyParser(&card); err != nil {
		panic(err)
	}

	database.DB.Model(&card).Updates(card)
	return c.Status(201).JSON(fiber.Map{
		"card": card,
	})
}

// 5- delete single deck
func DeleteCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		panic(err)
	}
	var card models.Card
	database.DB.Where("id = ?", id).Delete(&card)

	return c.Status(200).JSON(fiber.Map{
		"message": "deck recored delete successfully",
	})

}
