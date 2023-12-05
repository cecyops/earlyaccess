package handlers

import (
	"errors"

	"github.com/cecyops/earlyaccess/database"
	"github.com/cecyops/earlyaccess/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ListGebruikers(c *fiber.Ctx) error {
	gebruikers := []models.Gebruiker{}
	database.DB.Db.Find(&gebruikers)

	return c.Status(200).JSON(gebruikers)
}

func CreateGebruiker(c *fiber.Ctx) error {
	gebruiker := new(models.Gebruiker)
	if err := c.BodyParser(gebruiker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&gebruiker)

	return c.Status(200).JSON(gebruiker)
}

func ListSleutel(c *fiber.Ctx) error {
	sleutels := []models.Sleutel{}
	database.DB.Db.Find(&sleutels)

	return c.Status(200).JSON(sleutels)
}

func CreateSleutel(c *fiber.Ctx) error {
	sleutel := new(models.Sleutel)
	if err := c.BodyParser(sleutel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&sleutel)

	return c.Status(200).JSON(sleutel)
}

func AssignFirstAvailableSleutelToGebruiker(c *fiber.Ctx) error {
	gebruikerID := c.Params("gebruikerID")

	// Find the gebruiker
	var gebruiker models.Gebruiker
	if err := database.DB.Db.First(&gebruiker, gebruikerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Gebruiker not found",
		})
	}

	// Find the first available sleutel
	var sleutel models.Sleutel
	if err := database.DB.Db.Where("is_beschikbaar = ?", "ja").First(&sleutel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Er zijn geen sleutels meer beschikbaar",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Update sleutel as no longer available
	sleutel.IsBeschikbaar = "nee" // Assuming 'nee' means not available
	database.DB.Db.Save(&sleutel)

	// Update gebruiker with sleutel code
	gebruiker.Sleutel = sleutel.Code // Adjust this line according to your model structure
	gebruiker.Status = "Sleutel ontvangen"
	database.DB.Db.Save(&gebruiker)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"gebruiker": gebruiker,
		"sleutel":   sleutel,
	})
}
