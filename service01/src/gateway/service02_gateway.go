package gateway

import (
	"encoding/json"
	"model"
	"net/http"
	"net/url"
)

type Service02Gateway interface {
	Inquiry() (*model.Service02Model, error)
}

func NewHttpServiceGateway(baseURL *url.URL) Service02Gateway {
	return &httpService{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}
}

type httpService struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func (hs *httpService) Inquiry() (*model.Service02Model, error) {
	rel := &url.URL{Path: "/api/v1/mocks"}
	u := hs.baseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := hs.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result model.Service02Model
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}
