package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/goodsfeedback"
)

func (c *Client) DeleteGoodsFeedbackComment(ctx context.Context, businessId int64, body *models.DeleteGoodsFeedbackCommentRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsfeedback.DeleteGoodsFeedbackCommentEndpoint, utils.PathInt("businessId", businessId))
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

func (c *Client) GetGoodsFeedbackComments(ctx context.Context, businessId int64, params *models.GetGoodsFeedbackCommentsParams, body *models.GetGoodsFeedbackCommentsRequest) (*models.GetGoodsFeedbackCommentsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsfeedback.GetGoodsFeedbackCommentsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetGoodsFeedbackCommentsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetGoodsFeedbacks(ctx context.Context, businessId int64, params *models.GetGoodsFeedbacksParams, body *models.GetGoodsFeedbackRequest) (*models.GetGoodsFeedbackResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsfeedback.GetGoodsFeedbacksEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetGoodsFeedbackResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) SkipGoodsFeedbacksReaction(ctx context.Context, businessId int64, body *models.SkipGoodsFeedbackReactionRequest) (*models.EmptyApiResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsfeedback.SkipGoodsFeedbacksReactionEndpoint, utils.PathInt("businessId", businessId))
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

func (c *Client) UpdateGoodsFeedbackComment(ctx context.Context, businessId int64, body *models.UpdateGoodsFeedbackCommentRequest) (*models.UpdateGoodsFeedbackCommentResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsfeedback.UpdateGoodsFeedbackCommentEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.UpdateGoodsFeedbackCommentResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
