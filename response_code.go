package freshchat

type responseCode int

const (
	Success             responseCode = 200
	Accepted            responseCode = 202
	BadRequest          responseCode = 400
	Unauthenticated     responseCode = 401
	Forbidden           responseCode = 403
	NotFound            responseCode = 404
	TooManyRequests     responseCode = 429
	InternalServerError responseCode = 500
	ServiceUnavailable  responseCode = 503
)