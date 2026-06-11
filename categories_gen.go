package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils/categories"
)

func (c *Client) GetCategoriesMaxSaleQuantum(ctx context.Context, body *models.GetCategoriesMaxSaleQuantumRequest) (*models.GetCategoriesMaxSaleQuantumResponse, error) {
	requestURL := c.BaseURL + categories.GetCategoriesMaxSaleQuantumEndpoint
	var err error

	var reqBody io.Reader

	var jsonBody []byte
	jsonBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody = bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetCategoriesMaxSaleQuantumResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetCategoriesTree(ctx context.Context, body *models.GetCategoriesRequest) (*models.GetCategoriesResponse, error) {
	requestURL := c.BaseURL + categories.GetCategoriesTreeEndpoint
	var err error

	var reqBody io.Reader

	var jsonBody []byte
	jsonBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody = bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetCategoriesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
