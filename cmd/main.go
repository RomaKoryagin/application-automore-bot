package main

import (
	"fmt"
	"log"
	"os"

	"alex.com/application-bot/internal/application/factories"
	"alex.com/application-bot/internal/application/handlers"
	"alex.com/application-bot/internal/application/jobs"
	"alex.com/application-bot/internal/application/services"
	"alex.com/application-bot/internal/application/strategies"
	"alex.com/application-bot/internal/infrastructure/repositories"
	"alex.com/application-bot/internal/infrastructure/sqlite"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	executableFolderPath, _ := os.Getwd()

	fmt.Println(executableFolderPath)

	db := sqlite.Database{MainDirPath: executableFolderPath}

	db.Init()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	telegramGroupName := os.Getenv("GROUP_TELEGRAM_CHAT_NAME")

	/** repositories */
	userRepository := repositories.NewUserRepository(db.Connection)
	applicationRepository := repositories.NewApplicationRepository(db.Connection)
	/** end repositories */

	/** services */
	userService := services.NewUserService(userRepository)
	applicationService := services.NewApplicationService(applicationRepository, userRepository)
	/** end services */

	/** factories */
	factoryResolver := factories.NewFactoryResolver()
	/** end factories */

	/** strategies */
	startCommandStrategy := strategies.NewStartCommandStrategy(userService, applicationService)
	updateApplicationStrategy := strategies.NewUpdateApplicationStrategy(userService, applicationService, factoryResolver)
	websiteComamandStrategy := strategies.NewWebsiteCommandStrategy()
	aboutCommandStrategy := strategies.NewAboutCommandStrategy()
	noActiveApplicationStrategy := strategies.NewNoActiveApplicationStrategy()
	newApplicationStrategy := strategies.NewNewApplicationStrategy(userService, applicationService)
	menuCommandStrategy := strategies.NewMenuCommandStrategy()
	/** end strategies */

	/** resolvers */
	strategyResolver := strategies.NewStrategyResolver(applicationService)

	strategyResolver.AddStrategy(startCommandStrategy)
	strategyResolver.AddStrategy(updateApplicationStrategy)
	strategyResolver.AddStrategy(websiteComamandStrategy)
	strategyResolver.AddStrategy(menuCommandStrategy)
	strategyResolver.AddStrategy(aboutCommandStrategy)
	strategyResolver.AddStrategy(newApplicationStrategy)
	strategyResolver.AddStrategy(noActiveApplicationStrategy)
	/** end resolvers */

	/** services */
	telegramMessageService := services.NewTelegramMessageService(bot, strategyResolver)
	/** end services */

	/** handlers */
	applicationBotHandler := handlers.NewApplicationBotHandler(telegramMessageService)
	/** end handlers */

	jobs.NewSendTelegramApplicationJob(applicationService, bot, telegramGroupName).Execute()

	router := gin.Default()

	router.POST("/v1/application-bot/handle", applicationBotHandler.Handle)

	router.Run(fmt.Sprintf(":%s", os.Getenv("REST_API_PORT")))
}
