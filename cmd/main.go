package main

import (
	"fmt"
	"log"
	"os"

	"alex.com/application-bot/internal/application/factories"
	"alex.com/application-bot/internal/application/services"
	"alex.com/application-bot/internal/application/strategies"
	"alex.com/application-bot/internal/infrastructure/repositories"
	"alex.com/application-bot/internal/infrastructure/sqlite"
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

	/** repositories */
	userRepository := repositories.NewUserRepository(db.Connection)
	applicationRepository := repositories.NewApplicationRepository(db.Connection)
	/** end repositories */

	/** services */
	userService := services.NewUserService(userRepository)
	applicationService := services.NewApplicationService(applicationRepository)
	/** end services */

	/** factories */
	factoryResolver := factories.NewFactoryResolver()
	/** end factories */

	/** strategies */
	startCommandStrategy := strategies.NewStartCommandStrategy(userService, applicationService)
	updateApplicationStrategy := strategies.NewUpdateApplicationStrategy(userService, applicationService, factoryResolver)
	websiteComamandStrategy := strategies.NewWebsiteCommandStrategy()
	aboutCommandStrategy := strategies.NewAboutCommandStrategy()
	newApplicationStrategy := strategies.NewNewApplicationStrategy(userService, applicationService)
	/** end strategies */

	/** resolvers */
	strategyResolver := strategies.NewStrategyResolver()

	strategyResolver.AddStrategy(startCommandStrategy)
	strategyResolver.AddStrategy(updateApplicationStrategy)
	strategyResolver.AddStrategy(websiteComamandStrategy)
	strategyResolver.AddStrategy(aboutCommandStrategy)
	strategyResolver.AddStrategy(newApplicationStrategy)
	/** end resolvers */

	telegramMessageService := services.NewTelegramMessageService(bot, strategyResolver)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			go telegramMessageService.SendReplyMessage(update.Message.Chat.ID, update.Message.Text)
		}
	}
}
