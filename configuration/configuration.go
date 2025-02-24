package configuration

import (
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Templates         = template.Must(template.ParseGlob("views/*/*.html"))
	PortHTTPServer    = "9999"
	ServerDB          = "127.0.0.1"
	PortDB            = "3306"
	NameDB            = "tribegroups"
	UserDB            = "tribegroups"
	PasswordDB        = "tribegroups"
	ClientIDOAuth     = ""
	ClientSecretOauth = ""
	ClientCallBackURL = ""
)

func LoadOAuthData() {
	err := godotenv.Load("enviroment.env")
	if err != nil {
		log.Fatal(".env file failed to load!")
	}

	ClientIDOAuth = os.Getenv("CLIENT_ID")
	ClientSecretOauth = os.Getenv("CLIENT_SECRET")
	ClientCallBackURL = os.Getenv("CLIENT_CALLBACK_URL")

	if ClientIDOAuth == "" || ClientSecretOauth == "" || ClientCallBackURL == "" {
		log.Fatal("Environment variables (CLIENT_ID, CLIENT_SECRET, CLIENT_CALLBACK_URL) are required")
	}
}
