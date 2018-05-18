package model

import (
	"encoding/json"
	"errors"
)

type responseCode struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

type paginator struct {
	TotalCount int `json:"total_count"`
	Page       int
	Pagesize   int
}

type responseResult struct {
	responseCode
	Data bool
}

//StringToResult json convert to result
func StringToResult(str string) (bool, error) {

	var data responseResult
	json.Unmarshal([]byte(str), &data)

	if data.ErrNo != 0 {
		return false, errors.New(data.ErrMsg)
	}
	return data.Data, nil

}
