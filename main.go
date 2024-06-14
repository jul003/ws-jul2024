package main

import (
	"log"

	"github.com/jul003/ws-jul2024/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"


	"github.com/jul003/ws-jul2024/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jul003/ws-jul2024/docs"

)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/jul003
// @contact.email indra@ulbi.ac.id

// @host  https://ws-jul2024-bdc7cc35640a.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
