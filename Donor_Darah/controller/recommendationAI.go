package controller

import (
	"app/utils"
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

func RecommendationAI(c echo.Context) error {
	userInput := c.FormValue("Gol_Darah")

	// Meminta rekomendasi golongan darah yang cocok kepada AI
	client := openai.NewClient(os.Getenv("AI_TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda cukup memberitahu golongan darah mana yang dapat diterima oleh golongan darah: " + userInput,
				},
				// {
				// 	Role:    openai.ChatMessageRoleUser,
				// 	Content: "Saya butuh transfusi darah. Golongan darah mana yang bisa saya terima? Golongan darah saya adalah: " + userInput, // Menggunakan golongan darah dari pesanan sebagai pertanyaan
				// },
			},
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghubungi AI"))
	}

	// Mengambil jawaban dari AI
	recommendedBloodType := resp.Choices[0].Message.Content

	return c.JSON(http.StatusOK, utils.SuccessResponse("Rekomendasi: ", recommendedBloodType))
}
