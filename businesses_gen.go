package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/businesses"
)

func (c *Client) GetBusinessSettings(ctx context.Context, businessId int64) (*models.GetBusinessSettingsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(businesses.GetBusinessSettingsEndpoint, utils.PathInt("businessId", businessId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetBusinessSettingsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
