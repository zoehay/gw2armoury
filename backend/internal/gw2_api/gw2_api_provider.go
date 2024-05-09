package gw2api

import (
	"fmt"
	"io/ioutil"

	"encoding/json"
	"net/http"

	apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/api_models"
)

// getAllItems - on server init

// getCharacterInventory bags: Bag[]

func GetSomeItems() (*[]apimodels.ApiItem, error) {
	url := "https://api.guildwars2.com/v2/items?ids=24,68"

	// headers := http.Header{}
	// headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	// params := make(map[string]string)
	// params["ids"] = request.Ids

	// res, err := clients.Get(url, params, headers)
	// fmt.Print(res)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	bodyRaw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result []apimodels.ApiItem

	if err = json.Unmarshal(bodyRaw, &result); err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &result, nil
}
