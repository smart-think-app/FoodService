package common_model

type SuccessResponseDto struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
