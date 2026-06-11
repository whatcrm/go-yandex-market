package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/goodsquestions"
)

func (c *Client) GetGoodsQuestionAnswers(ctx context.Context, businessId int64, params *models.GetGoodsQuestionAnswersParams, body *models.GetAnswersRequest) (*models.GetAnswersResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsquestions.GetGoodsQuestionAnswersEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetAnswersResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetGoodsQuestions(ctx context.Context, businessId int64, params *models.GetGoodsQuestionsParams, body *models.GetQuestionsRequest) (*models.GetQuestionsResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsquestions.GetGoodsQuestionsEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.GetQuestionsResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) UpdateGoodsQuestionTextEntity(ctx context.Context, businessId int64, body *models.UpdateGoodsQuestionTextEntityRequest) (*models.UpdateGoodsQuestionTextEntityResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(goodsquestions.UpdateGoodsQuestionTextEntityEndpoint, utils.PathInt("businessId", businessId))
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

	var response models.UpdateGoodsQuestionTextEntityResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
