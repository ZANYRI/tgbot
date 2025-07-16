package main

import (
	"context"
	"log"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func ReqApi(message string) (string, error) {

	config, err := ReadEnv()

	if err != nil {
		log.Fatal("Error read env!")
	} 

	client := openai.NewClient(
	option.WithAPIKey(config.ApiKey),
	option.WithBaseURL("https://openrouter.ai/api/v1"),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Model: "moonshotai/kimi-k2:free",
	})
	if err != nil {
		panic(err.Error())
	}

	return chatCompletion.Choices[0].Message.Content, nil

}