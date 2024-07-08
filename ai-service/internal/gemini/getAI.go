package gemini

import (
	"armiya/ai-service/internal/config"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ConnectAI(config config.Config) (*genai.GenerativeModel, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}
	fmt.Println("API Key Loaded Successfully")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	modelName := "gemini-1.5-flash"

	model := client.GenerativeModel(modelName)

	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockOnlyHigh,
		},
	}

	model.GenerationConfig = genai.GenerationConfig{
		ResponseMIMEType: "application/json",
	}

	return model, nil
}
