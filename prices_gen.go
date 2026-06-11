package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/prices"
)

func (c *Client) GetDefaultPrices(ctx context.Context, businessId int64, params *models.GetDefaultPricesParams, body *models.GetDefaultPricesRequest) (*models.GetDefaultPricesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(prices.GetDefaultPricesEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetDefaultPricesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetPrices(ctx context.Context, campaignId int64, params *models.GetPricesParams) (*models.GetPricesResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(prices.GetPricesEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetPricesResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetPricesByOfferIds(ctx context.Context, campaignId int64, params *models.GetPricesByOfferIdsParams, body *models.GetPricesByOfferIdsRequest) (*models.GetPricesByOfferIdsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(prices.GetPricesByOfferIdsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetPricesByOfferIdsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateBusinessPrices(ctx context.Context, businessId int64, body *models.UpdateBusinessPricesRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(prices.UpdateBusinessPricesEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdatePrices(ctx context.Context, campaignId int64, body *models.UpdatePricesRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(prices.UpdatePricesEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.EmptyApiResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
