package services

import (
	"encoding/json"

	"github.com/yuki5155/go-llms/openai-llm/utils"
	"github.com/yuki5155/go-ocrs/prompts"
	"github.com/yuki5155/go-ocrs/schemas"
)

func NewReciptService(apiKey string, model string, reciptUrl string) (*schemas.StoreResponse, error) {

	config := utils.NewClientConfig(apiKey)
	client := utils.NewClient(config)
	reciptSchema := schemas.NewStoreSchema()
	schemaJSON, err := json.Marshal(reciptSchema)
	if err != nil {
		return nil, err
	}
	messages := []utils.Message{
		utils.NewMessageWithImage(
			reciptUrl,
			prompts.StoreSchemaPrompt,
		),
	}
	opts := utils.RequestOptions{
		Messages: messages,
		Schema:   schemaJSON,
	}
	res, err := client.SendRequestWithStructuredOutput(opts)
	if err != nil {

		return nil, err
	}
	imageAnalyze, err := utils.HandleResponse[schemas.StoreResponse](res)
	if err != nil {
		return nil, err
	}
	return imageAnalyze, nil
}
