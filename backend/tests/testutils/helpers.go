package testutils

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/zoehay/gw2armoury/backend/internal/api/models"
)

func PrintObject(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func UnmarshalResponse(w *httptest.ResponseRecorder) (map[string]interface{}, error) {
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func UnmarshalAccount(w *httptest.ResponseRecorder) (*models.Account, error) {
	var response *models.Account
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
