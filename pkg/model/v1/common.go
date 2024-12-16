package modelv1

import "github.com/centraldigital/cfw-sales-x-ordering-api/internal/constant"

type EmptyStruct struct{}

type SuccessResponse struct {
	Code    string                  `json:"code" default:"0000"`
	Message constant.SuccessMessage `json:"description"`
}

var (
	SuccessResponse_Created       = SuccessResponse{Message: constant.SuccessMessage_Created, Code: constant.SuccessCode}
	SuccessResponse_Updated       = SuccessResponse{Message: constant.SuccessMessage_Updated, Code: constant.SuccessCode}
	SuccessResponse_Deleted       = SuccessResponse{Message: constant.SuccessMessage_Deleted, Code: constant.SuccessCode}
	SuccessResponse_ResetPassword = SuccessResponse{Message: constant.SuccessMessage_ResetPassword, Code: constant.SuccessCode}
)
