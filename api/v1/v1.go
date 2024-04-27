package v1

const (
	StatusOk      = 200
	StatusCreated = 201

	StatusBadRequest       = 400
	StatusUnauthorized     = 401
	StatusForbidden        = 403
	StatusNotFound         = 404
	StatusMethodNotAllowed = 405
	StatusRequestTimeout   = 408
	StatusConflict         = 409
	StatusTooManyRequests  = 429

	StatusInternalServerError = 500
	StatusServiceUnavailable  = 503
	StatusGatewayTimeout      = 504

	StatusMongoError = 600
	StatusRedisError = 601
)
