package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Environment struct {
	FrontendUrl          string `envconfig:"FRONTEND_URL"`
	MetaClientID         string `envconfig:"META_CLIENT_ID"`
	MetaClientSecret     string `envconfig:"META_CLIENT_SECRET"`
	RedisAddr            string `envconfig:"REDIS_ADDR"`
	DatabaseUser         string `envconfig:"DATABASE_USER"`
	DatabasePort         string `envconfig:"DATABASE_PORT"`
	DatabasePass         string `envconfig:"DATABASE_PASSWORD"`
	DatabaseName         string `envconfig:"DATABASE_SCHEME"`
	DatabaseHost         string `envconfig:"DATABASE_HOST"`
	AccessSecretKey      string `envconfig:"ACCESS_SECRET_KEY"`
	WordpressAdminEmail  string `envconfig:"WORDPRESS_ADMIN_EMAIL"`
	GraphApiURL          string `envconfig:"GRAPH_API_URL"`
	SlackWebhookURL      string `envconfig:"SLACK_WEBHOOK_URL"`
	RodutKey             string `envconfig:"RODUT_KEY"`
	StripeEndpointSecret string `envconfig:"STRIPE_ENDPOINT_SECRET"`
	OpenAiApiKey         string `envconfig:"OPENAI_API_KEY"`
	SendgridApiKey       string `envconfig:"SENDGRID_API_KEY"`
	FromEmail            string `envconfig:"FROM_EMAIL"`
}

var Env Environment

func init() {
	if err := envconfig.Process("", &Env); err != nil {
		log.Fatal(err)
	}
}
