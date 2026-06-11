package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/stocks"
)

func (c *Client) GetStocks(ctx context.Context, campaignId int64, params *models.GetStocksParams, body *models.GetWarehouseStocksRequest) (*models.GetWarehouseStocksResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(stocks.GetStocksEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

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

	var response models.GetWarehouseStocksResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateStocks(ctx context.Context, campaignId int64, body *models.UpdateStocksRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(stocks.UpdateStocksEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	var reqBody io.Reader

	var jsonBody []byte
	jsonBody, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	reqBody = bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, "PUT", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
