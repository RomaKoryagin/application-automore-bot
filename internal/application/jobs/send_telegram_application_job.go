package jobs

import (
	"log"
	"time"

	"alex.com/application-bot/internal/application/builders"
	"alex.com/application-bot/internal/domain/entities"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ISendTelegramApplicationJobApplicationService interface {
	GetSubbmited() ([]*entities.Application, error)
	Update(appl *entities.Application) error
}

type SendTelegramApplicationJob struct {
	ApplicationRepository ISendTelegramApplicationJobApplicationService
	Bot                   *tgbotapi.BotAPI
	ChannelIdentifier     string
}

func (job SendTelegramApplicationJob) Execute() {
	go func() {
		for {
			appls, err := job.ApplicationRepository.GetSubbmited()
			if err != nil {
				log.Printf("Error while trying to get subbmited applications, more: %s", err)
				continue
			}

			for _, appl := range appls {
				msg := builders.NewApplicationMessageBuilder().
					ConfigureChannelName(job.ChannelIdentifier).
					ConfigureApplicationText(appl).
					AddOpenChatButton(appl.TelegramId).
					Build()

				_, err := job.Bot.Send(msg)
				if err != nil {
					log.Println(err)
					continue
				}
				appl.SendedToTelegram = true
				job.ApplicationRepository.Update(appl)
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func NewSendTelegramApplicationJob(
	applicationRepository ISendTelegramApplicationJobApplicationService,
	bot *tgbotapi.BotAPI,
	channelIdentifier string,
) *SendTelegramApplicationJob {
	return &SendTelegramApplicationJob{
		ApplicationRepository: applicationRepository,
		Bot:                   bot,
		ChannelIdentifier:     channelIdentifier,
	}
}
