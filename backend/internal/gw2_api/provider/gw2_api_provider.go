package gw2apiprovider

import (
	"fmt"
	"io/ioutil"

	gw2apimodels "github.com/zoehay/gw2armoury/backend/internal/gw2_api/models"

	"encoding/json"
	"net/http"
)

// getAllItems - on server init

// getCharacterInventory bags: Bag[]




func GetSomeItems() (*[]gw2apimodels.Gw2Item, error) {
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

	var result []gw2apimodels.Gw2Item

	if err = json.Unmarshal(bodyRaw, &result); err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &result, nil
}

// func (gw2apimodels.Gw2Item) ApiItemToItem() (models.Item) {
// 	return 
// }