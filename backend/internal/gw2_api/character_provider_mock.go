package gw2api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

type CharacterProviderMock struct{}

func (characterProvider *CharacterProviderMock) GetAllCharacters(apiKey string) ([]apimodels.APICharacter, error) {
	apiCharacters, err := characterProvider.ReadCharacterFromFile("/Users/zoehay/Projects/gw2armoury/backend/test_data/character_test_data.txt")

	if err != nil {
		return nil, fmt.Errorf("error reading from test data file: %s", err)
	}

	fmt.Println("PROVIDEER MOCK CHARADCTERS", apiCharacters[0].Name)

	return apiCharacters, nil
}

func (characterProvider *CharacterProviderMock) ReadCharacterFromFile(filepath string) ([]apimodels.APICharacter, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var characters []apimodels.APICharacter
	err = json.Unmarshal(content, &characters)
	if err != nil {
		return nil, err
	}

	return characters, nil
}
