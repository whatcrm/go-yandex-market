package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/orderdelivery"
)

func (c *Client) GetOrderBuyerInfo(ctx context.Context, campaignId int64, orderId int64) (*models.GetOrderBuyerInfoResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderdelivery.GetOrderBuyerInfoEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetOrderBuyerInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SetOrderDeliveryDate(ctx context.Context, campaignId int64, orderId int64, body *models.SetOrderDeliveryDateRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderdelivery.SetOrderDeliveryDateEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

func (c *Client) SetOrderDeliveryTrackCode(ctx context.Context, campaignId int64, orderId int64, body *models.SetOrderDeliveryTrackCodeRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderdelivery.SetOrderDeliveryTrackCodeEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

func (c *Client) UpdateOrderStorageLimit(ctx context.Context, campaignId int64, orderId int64, body *models.UpdateOrderStorageLimitRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderdelivery.UpdateOrderStorageLimitEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

func (c *Client) VerifyOrderEac(ctx context.Context, campaignId int64, orderId int64, body *models.VerifyOrderEacRequest) (*models.VerifyOrderEacResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderdelivery.VerifyOrderEacEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
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

	var response models.VerifyOrderEacResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
