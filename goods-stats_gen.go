package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/goodsstats"
)

func (c *Client) GetGoodsStats(ctx context.Context, campaignId int64, body *models.GetGoodsStatsRequest) (*models.GetGoodsStatsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsstats.GetGoodsStatsEndpoint, utils.PathInt("campaignId", campaignId))
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

	var response models.GetGoodsStatsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
