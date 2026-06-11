package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/bids"
)

func (c *Client) GetBidsInfoForBusiness(ctx context.Context, businessId int64, params *models.GetBidsInfoForBusinessParams, body *models.GetBidsInfoRequest) (*models.GetBidsInfoResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(bids.GetBidsInfoForBusinessEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetBidsInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetBidsRecommendations(ctx context.Context, businessId int64, body *models.GetBidsRecommendationsRequest) (*models.GetBidsRecommendationsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(bids.GetBidsRecommendationsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetBidsRecommendationsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) PutBidsForBusiness(ctx context.Context, businessId int64, body *models.PutSkuBidsRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(bids.PutBidsForBusinessEndpoint, utils.PathInt("businessId", businessId))
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

func (c *Client) PutBidsForCampaign(ctx context.Context, campaignId int64, body *models.PutSkuBidsRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(bids.PutBidsForCampaignEndpoint, utils.PathInt("campaignId", campaignId))
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
