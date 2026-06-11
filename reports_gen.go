package yandexmarket

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/whatcrm/go-yandex-market/models"
	"github.com/whatcrm/go-yandex-market/utils"
	"github.com/whatcrm/go-yandex-market/utils/reports"
)

func (c *Client) GenerateBannersStatisticsReport(ctx context.Context, params *models.GenerateBannersStatisticsReportParams, body *models.GenerateBannersStatisticsRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateBannersStatisticsReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateBarcodesReport(ctx context.Context, body *models.GenerateBarcodesReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateBarcodesReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateBoostConsolidatedReport(ctx context.Context, params *models.GenerateBoostConsolidatedReportParams, body *models.GenerateBoostConsolidatedRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateBoostConsolidatedReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateClosureDocumentsDetalizationReport(ctx context.Context, params *models.GenerateClosureDocumentsDetalizationReportParams, body *models.GenerateClosureDocumentsDetalizationRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateClosureDocumentsDetalizationReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateClosureDocumentsReport(ctx context.Context, body *models.GenerateClosureDocumentsRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateClosureDocumentsReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateCompetitorsPositionReport(ctx context.Context, params *models.GenerateCompetitorsPositionReportParams, body *models.GenerateCompetitorsPositionReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateCompetitorsPositionReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateGoodsFeedbackReport(ctx context.Context, params *models.GenerateGoodsFeedbackReportParams, body *models.GenerateGoodsFeedbackRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateGoodsFeedbackReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateGoodsMovementReport(ctx context.Context, params *models.GenerateGoodsMovementReportParams, body *models.GenerateGoodsMovementReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateGoodsMovementReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateGoodsPricesReport(ctx context.Context, params *models.GenerateGoodsPricesReportParams, body *models.GenerateGoodsPricesReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateGoodsPricesReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateGoodsRealizationReport(ctx context.Context, params *models.GenerateGoodsRealizationReportParams, body *models.GenerateGoodsRealizationReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateGoodsRealizationReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateGoodsTurnoverReport(ctx context.Context, params *models.GenerateGoodsTurnoverReportParams, body *models.GenerateGoodsTurnoverRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateGoodsTurnoverReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateJewelryFiscalReport(ctx context.Context, params *models.GenerateJewelryFiscalReportParams, body *models.GenerateJewelryFiscalReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateJewelryFiscalReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateKeyIndicatorsReport(ctx context.Context, params *models.GenerateKeyIndicatorsReportParams, body *models.GenerateKeyIndicatorsRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateKeyIndicatorsReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateMassOrderLabelsReport(ctx context.Context, params *models.GenerateMassOrderLabelsReportParams, body *models.GenerateMassOrderLabelsRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateMassOrderLabelsReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateSalesGeographyReport(ctx context.Context, params *models.GenerateSalesGeographyReportParams, body *models.GenerateSalesGeographyRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateSalesGeographyReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateShelfsStatisticsReport(ctx context.Context, params *models.GenerateShelfsStatisticsReportParams, body *models.GenerateShelfsStatisticsRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateShelfsStatisticsReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateShipmentListDocumentReport(ctx context.Context, body *models.GenerateShipmentListDocumentReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateShipmentListDocumentReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateShowsBoostReport(ctx context.Context, params *models.GenerateShowsBoostReportParams, body *models.GenerateShowsBoostRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateShowsBoostReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateShowsSalesReport(ctx context.Context, params *models.GenerateShowsSalesReportParams, body *models.GenerateShowsSalesReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateShowsSalesReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateStocksOnWarehousesReport(ctx context.Context, params *models.GenerateStocksOnWarehousesReportParams, body *models.GenerateStocksOnWarehousesReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateStocksOnWarehousesReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateUnitedMarketplaceServicesReport(ctx context.Context, params *models.GenerateUnitedMarketplaceServicesReportParams, body *models.GenerateUnitedMarketplaceServicesReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateUnitedMarketplaceServicesReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateUnitedNettingReport(ctx context.Context, params *models.GenerateUnitedNettingReportParams, body *models.GenerateUnitedNettingReportRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateUnitedNettingReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateUnitedOrdersReport(ctx context.Context, params *models.GenerateUnitedOrdersReportParams, body *models.GenerateUnitedOrdersRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateUnitedOrdersReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GenerateUnitedReturnsReport(ctx context.Context, params *models.GenerateUnitedReturnsReportParams, body *models.GenerateUnitedReturnsRequest) (*models.GenerateReportResponse, error) {
	requestURL := c.BaseURL + reports.GenerateUnitedReturnsReportEndpoint
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

	var response models.GenerateReportResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}

func (c *Client) GetReportInfo(ctx context.Context, reportId string) (*models.GetReportInfoResponse, error) {
	requestURL := c.BaseURL + utils.BuildPath(reports.GetReportInfoEndpoint, utils.PathStringParam("reportId", reportId))
	var err error

	var reqBody io.Reader

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, reqBody)
	if err != nil {
		return nil, err
	}

	var response models.GetReportInfoResponse
	if err = c.Send(req, &response); err != nil {
		return nil, err
	}
	return &response, nil

}
