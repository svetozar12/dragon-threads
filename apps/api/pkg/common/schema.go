package common

type CommonErrorSchema struct {
	Error string `json:"error"`
}

type CommonPaginationSchema struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
	HasNext    bool  `json:"hasNext"`
}
