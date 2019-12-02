package resp

import (
	"github.com/chaincode_learning_go/chaincode/token/domain/protos"
	"github.com/golang/protobuf/proto"
)

func ToResponseBytes(message proto.Message) []byte {
	result, err := proto.Marshal(message)
	if err != nil {
		return []byte("")
	}
	return result
}

func buildCommonError(errCode string, message string) *protos.CommonErrorResponse {
	return &protos.CommonErrorResponse{
		ErrorCode:    errCode,
		ErrorMessage: message,
	}
}

func BuildErrorCommonResponse(statusCode string, errCode string, message string) *protos.CommonResponse {
	return &protos.CommonResponse{
		StatusCode: statusCode,
		Error:      buildCommonError(errCode, message),
	}
}

func BuildSuccessCommonResponse(statusCode string) *protos.CommonResponse {
	return &protos.CommonResponse{
		StatusCode: statusCode,
		Error:      nil,
	}
}
