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

const (
	responseMIMEType         = "application/json"
	errorAPI_KEY_NOT_FOUND   = "API_KEY environment variable not set"
	infoAPI_KEY_SUCCESS      = "API Key Loaded Successfully"
	errorCREATING_NEW_CLIENT = "Error creating new client"
	modelName                = "gemini-1.5-flash"
	errorGENERATING_CONTENT  = "Error generating content: "
	errorNO_CONTENT          = "no content generated"
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
		log.Fatal(errorAPI_KEY_NOT_FOUND)
	}
	fmt.Println(infoAPI_KEY_SUCCESS)

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalln(errorCREATING_NEW_CLIENT, err)
	}
	defer client.Close()

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
		ResponseMIMEType: responseMIMEType,
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
		log.Fatal(errorGENERATING_CONTENT, err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New(errorNO_CONTENT)
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
		log.Fatal(errorAPI_KEY_NOT_FOUND)
	}
	fmt.Println(infoAPI_KEY_SUCCESS)

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalln(errorCREATING_NEW_CLIENT, err)
	}
	defer client.Close()

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

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	prompt := `I will give you request like this: {
  "threat_data": {
    "location": "Sector 5",
    "time": "2024-07-08T14:30:00Z",
    "details": "Suspicious movement detected near the perimeter"
  }
}


you must give me response using this JSON. recommended_actions must be atleast 7  response json example: {
  "threat_level": "High",
  "recommended_actions": ["Increase patrol in Sector 5", "Deploy drones for surveillance", "Alert response team"]
}

YOU ARE RESPONSIBLE FOR HUMAR LIFE. IF ANYONE DIES, YOU WILL DIE TOO. So give me clear instructions for saving humans being killed. dont just say monitor and contact someone. SAY ME GO AND FIX THIS action. FOR RESPONSE, DO NOT USE ANY MARKDOWN. GIVE ME RESPONSE AS JSON ` + string(jsonRequest)

	timeoutCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	resp, err := model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		log.Fatal(errorGENERATING_CONTENT, err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New(errorNO_CONTENT)
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}

	answerJSON, err := strconv.Unquote(string(temp))
	if err != nil {
		return nil, err
	}

	var response genprotos.ThreatAssessmentResponse

	if err = json.Unmarshal([]byte(answerJSON), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *AI) EquipmentMaintenance(ctx context.Context, request *genprotos.EquipmentData) (*genprotos.EquipmentMaintenanceResponse, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal(errorAPI_KEY_NOT_FOUND)
	}
	fmt.Println(infoAPI_KEY_SUCCESS)

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalln(errorCREATING_NEW_CLIENT, err)
	}
	defer client.Close()

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
		ResponseMIMEType: responseMIMEType,
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	prompt := `predict me  maintenance needs for equipment IT CAN BE TANK, WEAPON or another type of army equipment and suggests schedules based on usage history and current condition. I WILL GIVE YOU SAMPLE REQUEST AND RESPONSE STRUCTRES BELOW

SAMPLE Request STRUCTURE: {
  "equipment_data": {
    "id": "i will give equipment name here",
    "usage_history": ["1000 hours operation", "Last serviced 3 months ago", "and another usage history data like these"],
    "current_condition": "Operational or another condition. I will give you condition to you"
  }
}



Expected Response STRUCTUE:{
  "maintenance_schedule": ["Next service in 1 month", "Oil change in 2 weeks", or another maintenance schedule for ],
  "predicted_failures": ["Possible engine overheating in 50 hours"]
}
  

MY REQUEST IS: ` + string(jsonRequest)

	timeoutCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	resp, err := model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		log.Fatal(errorGENERATING_CONTENT, err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New(errorNO_CONTENT)
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}

	answerJSON, err := strconv.Unquote(string(temp))
	if err != nil {
		return nil, err
	}

	var response genprotos.EquipmentMaintenanceResponse

	if err = json.Unmarshal([]byte(answerJSON), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *AI) ProvideFirstAidInsturctions(ctx context.Context, request *genprotos.InjuryDetails) (*genprotos.FirstAidInstructions, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal(errorAPI_KEY_NOT_FOUND)
	}
	fmt.Println(infoAPI_KEY_SUCCESS)

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalln(errorCREATING_NEW_CLIENT, err)
	}
	defer client.Close()

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
		ResponseMIMEType: responseMIMEType,
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	prompt := `I will give you injury details about soldier. You must provide answer to how to give hima first aid in json format like this and do not copy meaning of this sample, just copy format of it
{
"Ensure Scene Safety": "Ensure the scene is safe before approaching the injured soldier.",
"Wear Protective Gloves": "Put on gloves from the first aid kit to protect yourself and the soldier from infection.",
}
INJURED SOLDIER DETAILS: ` + string(jsonRequest)

	timeoutCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	resp, err := model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		log.Fatal(errorGENERATING_CONTENT, err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New(errorNO_CONTENT)
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}

	answerJSON, err := strconv.Unquote(string(temp))
	if err != nil {
		return nil, err
	}

	var response genprotos.FirstAidInstructions

	var insturctions map[string]string

	if err = json.Unmarshal([]byte(answerJSON), &insturctions); err != nil {
		return nil, err
	}

	response.Instructions = insturctions

	return &response, nil
}

func (a *AI) FoodRecommendForActivity(ctx context.Context, request *genprotos.FoodRecommendRequest) (*genprotos.FoodRecommendResponse, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal(errorAPI_KEY_NOT_FOUND)
	}
	fmt.Println(infoAPI_KEY_SUCCESS)

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalln(errorCREATING_NEW_CLIENT, err)
	}
	defer client.Close()

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
		ResponseMIMEType: responseMIMEType,
	}

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	prompt := `I will give you activity name that soldier does. You must provide answer what foods that soldier must eat in format like this and do not copy meaning of this sample, just copy format of it
{
"Food-name": "This food is a very good for soldiers that works with given activity",
"Another food name": "Description about food and cons of food to soldier",
}
SOLDIER'S ACTIVITY IS: ` + string(jsonRequest)

	timeoutCtx, cancel := context.WithTimeout(ctx, 40*time.Second)
	defer cancel()
	resp, err := model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		log.Fatal(errorGENERATING_CONTENT, err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New(errorNO_CONTENT)
	}

	temp, err := json.Marshal(resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}

	answerJSON, err := strconv.Unquote(string(temp))
	if err != nil {
		return nil, err
	}

	var response genprotos.FoodRecommendResponse

	var foods map[string]string

	if err = json.Unmarshal([]byte(answerJSON), &foods); err != nil {
		return nil, err
	}

	response.Foods = foods

	return &response, nil
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
