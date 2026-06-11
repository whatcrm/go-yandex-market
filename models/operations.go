package models

const (
	OperationStatusTypeDONE       OperationStatusType = "DONE"
	OperationStatusTypeFAILED     OperationStatusType = "FAILED"
	OperationStatusTypeINPROGRESS OperationStatusType = "IN_PROGRESS"
)

type GetOperationsRequest struct {
	OperationIds []OperationId `json:"operationIds"`

	OperationType OperationType `json:"operationType"`
}

type GetOperationsResponse struct {
	Result *GetOperationsResultDTO `json:"result,omitempty"`

	Status ApiResponseStatusType `json:"status"`
}

type GetOperationsResultDTO struct {
	Operations []OperationResultDTO `json:"operations"`
}

type OperationDTO struct {
	Id string `json:"id"`

	Type OperationType `json:"type"`
}

type OperationId = string

type OperationResultDTO struct {
	Id string `json:"id"`

	Status OperationStatusType `json:"status"`

	Type OperationType `json:"type"`
}

type OperationStatusType string

type OperationType string

type QuestionsTextEntityOperationType string

type GetOperationsJSONRequestBody = GetOperationsRequest
