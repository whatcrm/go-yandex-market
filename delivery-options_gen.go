package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/deliveryoptions"
)

func (c *Client) GetDeliveryOptions(ctx context.Context, campaignId int64, body *models.GetDeliveryOptionsRequest) (*models.GetDeliveryOptionsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(deliveryoptions.GetDeliveryOptionsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetDeliveryOptionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetReturnDeliveryOptions(ctx context.Context, campaignId int64, body *models.GetReturnDeliveryOptionsRequest) (*models.GetReturnDeliveryOptionsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(deliveryoptions.GetReturnDeliveryOptionsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetReturnDeliveryOptionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
