package main

import (
	"app/config"
	"app/routes"
	// "bufio"
	// "context"
	// "fmt"
	// "log"
	// "os"

	// "github.com/joho/godotenv"
	// "github.com/sashabaranov/go-openai"
)

func main() {
	config.ConnectDB()

	// loadEnv()
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// line := scanner.Text()
	// fmt.Println("captured:", line)
	// client := openai.NewClient(os.Getenv("AI_TOKEN"))
	// resp, err := client.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role: openai.ChatMessageRoleUser,
	// 				Content: line,
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Printf("chatCompletion error: %v\n", err)
	// 	return
	// }

	// fmt.Println(resp.Choices[0].Message.Content)

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}

// func loadEnv(){
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }
