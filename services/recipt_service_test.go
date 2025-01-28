package services_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/yuki5155/go-ocrs/services"
)

func TestNewReciptService(t *testing.T) {
	api_key := os.Getenv("OPENAI_API_KEY")
	if api_key == "" {
		t.Errorf("OPENAI_API_KEY must be set")
	}
	// レシートのurlを入れる
	sampleUrl := ""
	reciptService, err := services.NewReciptService(api_key, "model", sampleUrl)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if reciptService == nil {
		t.Errorf("reciptService is nil")
	}
	fmt.Println(reciptService)
}
