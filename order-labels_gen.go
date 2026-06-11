package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/orderlabels"
)

func (c *Client) GenerateOrderLabel(ctx context.Context, campaignId int64, orderId int64, shipmentId int64, boxId int64, params *models.GenerateOrderLabelParams) error {
	requestURL := c.BaseURL + utils.BuildPath(orderlabels.GenerateOrderLabelEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId), utils.PathInt("shipmentId", shipmentId), utils.PathInt("boxId", boxId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return err
	}

	if err = c.Send(req, nil); err != nil {
		return err
	}
	return nil

}

func (c *Client) GenerateOrderLabels(ctx context.Context, campaignId int64, orderId int64, params *models.GenerateOrderLabelsParams) error {
	requestURL := c.BaseURL + utils.BuildPath(orderlabels.GenerateOrderLabelsEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return err
	}

	if err = c.Send(req, nil); err != nil {
		return err
	}
	return nil

}

func (c *Client) GetOrderLabelsData(ctx context.Context, campaignId int64, orderId int64) (*models.GetOrderLabelsDataResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(orderlabels.GetOrderLabelsDataEndpoint, utils.PathInt("campaignId", campaignId), utils.PathInt("orderId", orderId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetOrderLabelsDataResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
