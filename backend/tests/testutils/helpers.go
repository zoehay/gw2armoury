package testutils

import "encoding/json"

func PrintObject(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
