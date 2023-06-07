package utils

import (
	"github.com/tkanos/gonfig"
)

type Configurations struct {
	DatabaseConnectionString string `env:"DB_CONNECTION_STRING"`
	OutboxCycleTime          int32
	LinkValidityTime         int32
	SendgridApiKey           string `env:"SENDGRID_API_KEY"`
	SenderEmailAddress       string
	SenderTitle              string
	DashboardAddress         string `env:"DASHBOARD_ADDRESS"`
	RegisterSub              string
	LoginSub                 string
	LoginTimeOut             int32
	TokenValidityTime        int32
	ProjectID                string `env:"GCP_PROJECT_ID"`
	AckDeadline              int
	FireBaseCredentialFile   string
	ServiceName              string
	MaxConcurrency           int
	ResentMailTimeOut        string
	MaxSecurityEvents        int
	LogLevel                 string
	RedirectUrl              string
	VerifyUrl                string
	StripeSecretKey          string `env:"STRIPE_SECRET_KEY"`
	WebLink                  string
	SystemLink               string
	FeedbackLink             string
	PolicyPathFile           string
	BillingInternalBaseUrl   string
	TokenGenKey              string `env:"TOKEN_GEN_KEY"`
	BundleVerificationKey    string `env:"BUNDLE_VERIFICATION_KEY"`
	FeedbackReceiverEmail    string
	RedisAddress             string `env:"REDIS_ADDRESS"`
	RedisPassword            string `env:"REDIS_PASSWORD"`
	RedisDatabase            int    `env:"REDIS_DATABASE"`
	SlackToken               string `env:"SLACK_APP_TOKEN"`
	SlackChannel             string `env:"SLACK_CHANNEL"`
}

func LoadConfigurations(path string) (*Configurations, error) {
	c := Configurations{}

	err := gonfig.GetConf(path, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
