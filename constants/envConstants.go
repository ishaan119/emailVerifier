package constants

import (
	"os"

	en "github.com/caarlos0/env"
)

//Env contains all the environment variables
var Env envVars

func init() {
	err := en.Parse(&Env)
	if err != nil {
		os.Exit(1)
	}
}

type envVars struct {
	//newrelic config
	NewrelicAppKey  string `env:"NEWRELIC_APP_KEY" envDefault:""`
	NewrelicAppName string `env:"NEWRELIC_APP_NAME" envDefault:"newsletter-engine"`

	//apis
	Port                 string `env:"PORT" envDefault:"9494"`
	NewsLetterEngineHost string `env:"NEWSLETTER_ENGINE_SERVICE_HOST" envDefault:"127.0.0.1"`
	NewsLetterEnginePort string `env:"NEWSLETTER_ENGINE_SERVICE_PORT" envDefault:"9494"`

	//aerospike
	AerospikeUrls []string `env:"AEROSPIKE_HOST_URL" envDefault:"localhost"`

	//statsd
	StatsdHost string `env:"F1COUNTER_SERVICE_HOST" envDefault:"f1counter.cfstage.com"`
	StatsdPort string `env:"F1COUNTER_SERVICE_PORT_STATSD" envDefault:"30519"`

	//postgres configs
	PostgresUserName     string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresUserPassword string `env:"POSTGRES_PASSWORD" envDefault:"Postgres@1"`
	PostgresDBName       string `env:"POSTGRES_DB" envDefault:"postgres"`
	PostgresHost         string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort         string `env:"POSTGRES_PORT" envDefault:"5432"`

	//proxybox configs
	ProxyboxHost string `env:"PROXYBOX_SERVICE_HOST" envDefault:"localhost"`
	ProxyboxPort string `env:"PROXYBOX_SERVICE_PORT" envDefault:"3050"`

	//crowdfire
	CrowdfireAPIKey string `env:"CROWDFIRE_API_KEY" envDefault:"awe23qewdawdasdads"`
	CrowdfireHost   string `env:"JU_INTERNAL_SERVICE_HOST" envDefault:"localhost"`
	CrowdfirePort   string `env:"JU_INTERNAL_SERVICE_PORT" envDefault:"9099"`

	//property-engine configs
	PropertyEngineHost string `env:"PROPERTY_ENGINE_SERVICE_HOST" envDefault:"localhost"`
	PropertyEnginePort string `env:"PROPERTY_ENGINE_SERVICE_PORT" envDefault:"3052"`

	//chaukidaar configs
	ChaukidaarBaseURL string `env:"CHAUKIDAAR_BASE_URL" envDefault:"http://api.cfstage.com"`

	//aws configs
	SqsEndPoint string `env:"SQS_END_POINT" envDefault:""`
	AwsRegion   string `env:"AWS_REGION" envDefault:"us-east-1"`

	//later service config
	LaterHost string `env:"LATER_SERVICE_HOST" envDefault:"localhost"`
	LaterPort string `env:"LATER_SERVICE_PORT" envDefault:"2000"`

	//sendgrid configs
	SendGridAPIKey string `env:"SENDGRID_API_KEY" envDefault:"dummyKey"`

	//template directory path
	TemplateDIR string `env:"TEMPLATE_DIR" envDefault:"templates"`

	// encryption-decryption key for crypto util
	UserDataEncryptionKey string `env:"USER_DATA_ENCRYPTION_KEY" envDefault:"wruvacr74ufrahup"`

	//FirehoseNewsletterActivityStream Email Activity Stream Name
	FirehoseNewsletterActivityStream string `env:"FIREHOSE_NEWSLETTER_ACTIVITY_STREAM" envDefault:"email_skill_newsletter_created"`
	//FirehoseSubscriberActivityStream Email Activity Stream Name
	FirehoseSubscriberActivityStream string `env:"FIREHOSE_SUBSCRIBER_ACTIVITY_STREAM" envDefault:"email_skill_subscribers"`
	//FirehoseEmailEventStream Email Activity Stream Name
	FirehoseEmailEventStream string `env:"FIREHOSE_EMAIL_EVENT_STREAM" envDefault:"email_skill_newsletter_event"`

	//content-engine configs
	ContentEngineHost string `env:"CONTENT_RECOMMENDATION_SERVICE_HOST" envDefault:"localhost"`
	ContentEnginePort string `env:"CONTENT_RECOMMENDATION_SERVICE_PORT" envDefault:"8000"`

	//is live server
	IsLiveServer bool `env:"IS_LIVE_SERVER" envDefault:"false"`

	MixpanelToken string `env:"MIXPANEL_TOKEN" envDefault:"dummy"`
}
