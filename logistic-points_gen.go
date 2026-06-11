package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/logisticpoints"
)

func (c *Client) GetLogisticPoints(ctx context.Context, businessId int64, params *models.GetLogisticPointsParams) (*models.GetLogisticPointsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(logisticpoints.GetLogisticPointsEndpoint, utils.PathInt("businessId", businessId))
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetLogisticPointsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
