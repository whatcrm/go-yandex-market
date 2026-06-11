package yandexmarket

import (
	"context"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/campaigns"
)

func (c *Client) GetCampaign(ctx context.Context, campaignId int64) (*models.GetCampaignResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(campaigns.GetCampaignEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetCampaignResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetCampaignSettings(ctx context.Context, campaignId int64) (*models.GetCampaignSettingsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(campaigns.GetCampaignSettingsEndpoint, utils.PathInt("campaignId", campaignId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetCampaignSettingsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetCampaigns(ctx context.Context, params *models.GetCampaignsParams) (*models.GetCampaignsResponse, error) {
	requestURL := c.BaseURL + campaigns.GetCampaignsEndpoint
	var err error

	requestURL, err = mergeQuery(requestURL, params)
	if err != nil {
		return nil, err
	}

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetCampaignsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
