package gemini

import (
	"armiya/ai-service/genprotos"
	"armiya/ai-service/internal/config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/k0kubun/pp"
	"google.golang.org/api/option"
)

type (
	AI struct {
		model *genai.GenerativeModel
	}
)

func New(config *config.Config) (*AI, error) {

	model, err := ConnectAI(*config)
	if err != nil {
		return nil, err
	}

	return &AI{
		model: model,
	}, nil
}

func (a *AI) GetEquipmentInfo(ctx context.Context, request *genprotos.EquipmentRequestAI) (*genprotos.EquipmentAI, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}
	fmt.Println("API Key Loaded Successfully")

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
	prompt := `I will give you a one sentence. it will be about any military equipment like tanks, weapons or other things.
You must extract it from sentence and write answer in this json format. do not double unescape json please:
{
    "name": "you will write extracted equipment name here",
    "description": "you will write description about this equipment",
    "ammo_type": "You will write which ammo type is used for this equipment. If this equipment dont uses ammo, write here "nil"",
    "usage_tutorial": [
        {
            "step_number": "1",
            "step": "You will write here tutorial step by step"
        },
        {
            "step_number": "2",
            "step": "You will write here titorial step by step"
        }
    ]
}

IN USAGE TUTORIAL, WRITE EACH STEP like {1: "go into tank or give weapon hands on please clarify each steps so you must say which button must b prrssed after which button in real life. WRITE AT LEAST 10 steps", 2: "reload guns"} SO LETS GO: SENTENCE IS ` + request.Prompt

	timeoutCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	resp, err := model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("no content generated")
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}

	unescapedData, err := strconv.Unquote(string(temp))
	if err != nil {
		pp.Println(err)
	}

	var carbine map[string]interface{}
	err = json.Unmarshal([]byte(unescapedData), &carbine)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	equipmentAI := convertMapToEquipmentAI(carbine)
	return equipmentAI, nil
}

func (a *AI) AssessThreat(ctx context.Context, request *genprotos.ThreatData) (*genprotos.ThreatAssessmentResponse, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable not set")
	}
	fmt.Println("API Key Loaded Successfully")

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
	prompt := `I will give you request with this JSON format: `

	timeoutCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	resp, err := model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("Error generating content: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("no content generated")
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}

	unescapedData, err := strconv.Unquote(string(temp))
	if err != nil {
		pp.Println(err)
	}

	var carbine map[string]interface{}
	err = json.Unmarshal([]byte(unescapedData), &carbine)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	equipmentAI := convertMapToEquipmentAI(carbine)
	pp.Println("ERROR YOQ", equipmentAI)
	return equipmentAI, nil
}

func convertMapToEquipmentAI(carbine map[string]interface{}) *genprotos.EquipmentAI {
	equipmentAI := &genprotos.EquipmentAI{
		Name:        carbine["name"].(string),
		Description: carbine["description"].(string),
		AmmoType:    carbine["ammo_type"].(string),
	}

	usageTutorial := carbine["usage_tutorial"].([]interface{})
	for _, step := range usageTutorial {
		stepMap := step.(map[string]interface{})
		equipmentAI.UsageTutorial = append(equipmentAI.UsageTutorial, &genprotos.UsageTutorialStepAI{
			StepNumber: stepMap["step_number"].(string),
			Step:       stepMap["step"].(string),
		})
	}

	return equipmentAI
}
