package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/orderbusinessinformation"
)

func (c *Client) GetOrderBusinessBuyerInfo(ctx context.Context, campaignId int64, orderId int64) (*models.GetBusinessBuyerInfoResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderbusinessinformation.GetOrderBusinessBuyerInfoEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetBusinessBuyerInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetOrderBusinessDocumentsInfo(ctx context.Context, campaignId int64, orderId int64) (*models.GetBusinessDocumentsInfoResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderbusinessinformation.GetOrderBusinessDocumentsInfoEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetBusinessDocumentsInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
