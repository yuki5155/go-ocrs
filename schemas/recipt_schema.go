package schemas

import (
	"github.com/yuki5155/go-llms/openai-llm/schema"
)

type StoreSchema struct {
	Name   string            `json:"name"`   // スキーマ名
	Schema schema.BaseSchema `json:"schema"` // スキーマ構造
}

// StoreResponse は店舗情報のレスポンスを定義します
type StoreResponse struct {
	StoreName string      `json:"store_name"` // 店舗名
	Date      string      `json:"date"`       // 日付 (yyyy-mm-dd)
	Items     []StoreItem `json:"items"`      // 品物リスト
}

// StoreItem は品物情報を表す構造体です
type StoreItem struct {
	Name  string `json:"name"`  // 品物名
	Price int    `json:"price"` // 価格(税込)
}

func NewStoreSchema() *StoreSchema {
	falseValue := false
	return &StoreSchema{
		Name: "store_response",
		Schema: schema.BaseSchema{
			Type: "object",
			Properties: map[string]schema.SchemaProperty{
				"store_name": {
					Type:        "string",
					Description: "Name of the store",
				},
				"date": {
					Type:        "string",
					Description: "Date in yyyy-mm-dd format",
				},
				"items": {
					Type:        "array",
					Description: "List of items available in the store",
					Items: &schema.SchemaProperty{
						Type: "object",
						Properties: map[string]schema.SchemaProperty{
							"name": {
								Type:        "string",
								Description: "Name of the item",
							},
							"price": {
								Type:        "integer",
								Description: "Price of the item",
							},
						},
						Required: []string{"name", "price"},
					},
				},
			},
			Required:             []string{"items"},
			AdditionalProperties: &falseValue,
		},
	}
}
