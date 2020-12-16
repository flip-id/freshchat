package freshchat_client

import (
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

var (
	baseUrl        string
	sendMessageUrl string
	apiToken       string
	namespace      string
	timeout        int
	client         *resty.Client
)

func init() {
	loadEnv()

	baseUrl = os.Getenv("FRESHCHAT_URL_BASE")
	sendMessageUrl = os.Getenv("FRESHCHAT_URL_SEND_MESSAGE")
	apiToken = os.Getenv("FRESHCHAT_API_TOKEN")
	namespace = os.Getenv("FRESHCHAT_NAMESPACE")

	timeout, _ = strconv.Atoi(os.Getenv("FRESHCHAT_TIMEOUT_IN_SECOND"))
	client = resty.New().
		SetTimeout(time.Second * time.Duration(timeout)).
		SetAuthToken(apiToken)
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		_ = godotenv.Load("./../../.env")
	}
}
