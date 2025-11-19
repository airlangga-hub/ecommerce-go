package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)


type AppConfig struct {
	ServerPort 			string
	Dsn 				string
	AppSecret 			string
	MyPhoneNumber 		string
	TwilioPhoneNumber 	string
	TwilioAccountSid 	string
	TwilioAuthToken 	string
}


func SetupEnv() (AppConfig, error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("http port env variable not found")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		return AppConfig{}, errors.New("dsn env variable not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if appSecret == "" {
		return AppConfig{}, errors.New("app secret env variable not found")
	}

	myPhoneNum := os.Getenv("MY_PHONE_NUMBER")
	if myPhoneNum == "" {
		return AppConfig{}, errors.New("my phone number env variable not found")
	}

	twilioPhoneNum := os.Getenv("TWILIO_PHONE_NUMBER")
	if twilioPhoneNum == "" {
		return AppConfig{}, errors.New("twilio phone number env variable not found")
	}

	twilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	if twilioAccountSid == "" {
		return AppConfig{}, errors.New("twilio account sid env variable not found")
	}

	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if twilioAuthToken == "" {
		return AppConfig{}, errors.New("twilio auth token env variable not found")
	}

	return AppConfig{
		ServerPort: httpPort,
		Dsn: dsn,
		AppSecret: appSecret,
		MyPhoneNumber: myPhoneNum,
		TwilioPhoneNumber: twilioPhoneNum,
		TwilioAccountSid: twilioAccountSid,
		TwilioAuthToken: twilioAuthToken,
		}, nil
}