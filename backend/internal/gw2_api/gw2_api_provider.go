package gw2api

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/zoehay/gw2armoury/backend/internal/clients"
	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

// getAllItems - on server init
// 		getAllItemIds
//		getItems?ids={batch of item ids} load items into db in batches

// getCharacterInventory bags: Bag[]

const baseUrl = "https://api.guildwars2.com/v2/"

func GetSomeItems(ids string) ([]apimodels.ApiItem, error) {
	// url := "https://api.guildwars2.com/v2/items?ids=24,68"
	// url := baseUrl + "items?ids=" + ids

	headers := http.Header{}
	// headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	params := make(map[string]string)
	params["ids"] = ids

	res, err := clients.Get(baseUrl, params, headers)

	// res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("http.Get error: %s", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll error: %s", err)
	}

	var result []apimodels.ApiItem

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal error: %s", err)
	}

	return result, nil
}
