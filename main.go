package main

import (
	"fmt"

	tgBotAPi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string `mapstructure:"tgToken"`
	GptToken      string `mapstructure:"gptToken"`
	Preamble      string `mapstructure:"preamble"`
}

func LoadConfig(path string) (c Config, err error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)     // path to look for the config file in

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}

func sendChatGPT(apiKey string) {

}

func main() {
 config, err := LoadConfig("./")

 if err != nil{
	panic(fmt.Errorf("fatal eror with the config.yaml: %w", err))
 }

 apiKey := config.GptToken
 fmt.Println(apiKey)

 tgBotAPi.NewBotAPI()

}