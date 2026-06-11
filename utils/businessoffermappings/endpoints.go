package businessoffermappings

const (
	AddOffersToArchiveEndpoint      = "/v2/businesses/{businessId}/offer-mappings/archive"
	DeleteOffersEndpoint            = "/v2/businesses/{businessId}/offer-mappings/delete"
	DeleteOffersFromArchiveEndpoint = "/v2/businesses/{businessId}/offer-mappings/unarchive"
	GenerateOfferBarcodesEndpoint   = "/v1/businesses/{businessId}/offer-mappings/barcodes/generate"
	GetOfferMappingsEndpoint        = "/v2/businesses/{businessId}/offer-mappings"
	UpdateOfferMappingsEndpoint     = "/v2/businesses/{businessId}/offer-mappings/update"
)
