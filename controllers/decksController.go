package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/flashcardApi/database"
	"github.com/zanamira43/flashcardApi/models"
)

// 1- get all decks
func GetALLDecks(c *fiber.Ctx) error {
	var decks []models.Deck
	database.DB.Preload("Card").Find(&decks)
	return c.Status(200).JSON(decks)
}

// 2- get create new deck
func CreateNewDeck(c *fiber.Ctx) error {
	var deck models.Deck

	if err := c.BodyParser(&deck); err != nil {
		return err
	}

	database.DB.Create(&deck)

	return c.Status(201).JSON(deck)

}

// 3- get single deck
func GetDeck(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		panic(err)
	}
	var deck models.Deck
	database.DB.Preload("Card").Where("id = ?", id).First(&deck)

	return c.Status(200).JSON(deck)

}

// 4- update single deck
func UpdateDeck(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		panic(err)
	}

	deck := models.Deck{
		ID: uint(id),
	}

	if err := c.BodyParser(&deck); err != nil {
		panic(err)
	}

	database.DB.Model(&deck).Updates(deck)
	return c.Status(201).JSON(deck)
}

// 5- delete single deck
func DeleteDeck(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		panic(err)
	}
	var deck models.Deck
	database.DB.Where("id = ?", id).Delete(&deck)

	return c.Status(200).JSON(fiber.Map{
		"message": "deck recored delete successfully",
	})

}
