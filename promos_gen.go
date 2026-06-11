package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/promos"
)

func (c *Client) DeletePromoOffers(ctx context.Context, businessId int64, body *models.DeletePromoOffersRequest) (*models.DeletePromoOffersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(promos.DeletePromoOffersEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.DeletePromoOffersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetPromoOffers(ctx context.Context, businessId int64, params *models.GetPromoOffersParams, body *models.GetPromoOffersRequest) (*models.GetPromoOffersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(promos.GetPromoOffersEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetPromoOffersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetPromos(ctx context.Context, businessId int64, body *models.GetPromosRequest) (*models.GetPromosResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(promos.GetPromosEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetPromosResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdatePromoOffers(ctx context.Context, businessId int64, body *models.UpdatePromoOffersRequest) (*models.UpdatePromoOffersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(promos.UpdatePromoOffersEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.UpdatePromoOffersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
