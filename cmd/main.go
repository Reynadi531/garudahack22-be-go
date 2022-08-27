package main

import (
	"flag"
	"gh22-go/config"
	"gh22-go/internal/routes"
	"gh22-go/internal/validations"
	"gh22-go/pkg/database"
	"gh22-go/pkg/middleware"
	"gh22-go/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	DevMode = flag.Bool("dev", false, "dev mode")
)

func init() {
	flag.Parse()
	config.InitConfig(*DevMode)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if *DevMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	database.InitDatabase()
	database.AutoMigrate(database.DB)
}

func main() {
	app := fiber.New()

	validations.InitValidations()

	middleware.RegisterMiddleware(app)

	routes.RegisterAuthRoutes(app, database.DB)
	routes.RegisterUserRoutes(app, database.DB)

	if *DevMode {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
