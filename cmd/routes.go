package main

import (
	"github.com/cecyops/earlyaccess/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/gebruikers", handlers.ListGebruikers)

	app.Post("/gebruiker", handlers.CreateGebruiker)

	app.Get("/sleutels", handlers.ListSleutel)

	app.Post("/sleutel", handlers.CreateSleutel)

	app.Put("/gebruiker/:gebruikerID/sleutel/assign", handlers.AssignFirstAvailableSleutelToGebruiker)

}
