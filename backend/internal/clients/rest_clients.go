package clients

import "net/http"

func Get(url string, params map[string]string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()
	
	req.Header = headers
	req.Header.Add("Content-Type", `application/json;charset=utf-8`)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
	
}