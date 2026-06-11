package yandexmarket

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/whatcrm/go-yandex-market/utils"
)

const DefaultBaseURL = "https://api.partner.market.yandex.ru"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("api key is required")
	}

	return &Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	if req.Header.Get("Content-Type") == "" && req.Method != http.MethodDelete && req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Api-Key", c.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-yandex-market-sdk")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return utils.ParseAPIError(resp)
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err = io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
