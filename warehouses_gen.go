package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/warehouses"
)

func (c *Client) GetFulfillmentWarehouses(ctx context.Context, params *models.GetFulfillmentWarehousesParams) (*models.GetFulfillmentWarehousesResponse, error) {
	requestURL := c.BaseURL + warehouses.GetFulfillmentWarehousesEndpoint
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetFulfillmentWarehousesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetPagedWarehouses(ctx context.Context, businessId int64, params *models.GetPagedWarehousesParams, body *models.GetPagedWarehousesRequest) (*models.GetPagedWarehousesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(warehouses.GetPagedWarehousesEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetPagedWarehousesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetWarehouses(ctx context.Context, businessId int64) (*models.GetWarehousesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(warehouses.GetWarehousesEndpoint, utils.PathInt("businessId", businessId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetWarehousesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateWarehouseStatus(ctx context.Context, campaignId int64, body *models.UpdateWarehouseStatusRequest) (*models.UpdateWarehouseStatusResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(warehouses.UpdateWarehouseStatusEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.UpdateWarehouseStatusResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
