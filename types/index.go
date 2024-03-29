package types

type Pagination struct {
	NextCursor string
	PageSize   int
}

type ApiResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
type ApiError struct {
	StatusCode int      `json:"statusCode"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
}

type PaginatedDataOutput struct {
	Data       interface{} `json:"data"`
	NextCursor string      `json:"nextCursor"`
	HasMore    bool        `json:"hasMore"`
	Total      int         `json:"total"`
}
