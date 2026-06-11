package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/ratings"
)

func (c *Client) GetQualityRatingDetails(ctx context.Context, campaignId int64) (*models.GetQualityRatingDetailsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(ratings.GetQualityRatingDetailsEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetQualityRatingDetailsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetQualityRatings(ctx context.Context, businessId int64, body *models.GetQualityRatingRequest) (*models.GetQualityRatingResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(ratings.GetQualityRatingsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetQualityRatingResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
