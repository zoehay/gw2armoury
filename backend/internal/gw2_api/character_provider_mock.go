package gw2api

import (
	"encoding/json"
	"fmt"
	"os"

	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

type CharacterProviderMock struct{}

func (characterProvider *CharacterProviderMock) GetAllCharacters(apiKey string) ([]apimodels.APICharacter, error) {
	apiCharacters, err := characterProvider.ReadCharactersFromFile("../test_data/character_test_data.txt")

	if err != nil {
		return nil, fmt.Errorf("error reading from test data file: %s", err)
	}

	return apiCharacters, nil
}

func (characterProvider *CharacterProviderMock) ReadCharactersFromFile(filepath string) ([]apimodels.APICharacter, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var characters []apimodels.APICharacter
	err = json.Unmarshal(content, &characters)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return characters, nil
}
