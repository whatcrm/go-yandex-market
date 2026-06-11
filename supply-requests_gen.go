package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/supplyrequests"
)

func (c *Client) GetSupplyRequestDocuments(ctx context.Context, campaignId int64, body *models.GetSupplyRequestDocumentsRequest) (*models.GetSupplyRequestDocumentsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(supplyrequests.GetSupplyRequestDocumentsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetSupplyRequestDocumentsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetSupplyRequestItems(ctx context.Context, campaignId int64, params *models.GetSupplyRequestItemsParams, body *models.GetSupplyRequestItemsRequest) (*models.GetSupplyRequestItemsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(supplyrequests.GetSupplyRequestItemsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetSupplyRequestItemsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetSupplyRequests(ctx context.Context, campaignId int64, params *models.GetSupplyRequestsParams, body *models.GetSupplyRequestsRequest) (*models.GetSupplyRequestsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(supplyrequests.GetSupplyRequestsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetSupplyRequestsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
