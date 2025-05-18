package enums

type StrategyType string

const (
	Start               StrategyType = "Start"
	About               StrategyType = "About"
	UpdateApplication   StrategyType = "UpdateApplications"
	WebsiteLink         StrategyType = "WebsiteLink"
	NewApplication      StrategyType = "NewApplication"
	Error               StrategyType = "Error"
	NoActiveApplication StrategyType = "NoActiveApplication"
	ShowMenu            StrategyType = "ShowMenu"
	JapanWheelWarning   StrategyType = "JapanWheelWarning"
	CountryResolving    StrategyType = "CountryResolving"
	CountryReturn       StrategyType = "CountryReturn"
)
