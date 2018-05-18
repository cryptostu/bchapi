package model

type responseCode struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

type paginator struct {
	TotalCount int `json:"total_count"`
	Page       int
	Pagesize   int
}
