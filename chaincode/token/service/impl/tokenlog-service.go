package impl

import (
	"github.com/chaincode_learning_go/chaincode/token/consts"
	"github.com/chaincode_learning_go/chaincode/token/domain/protos"
	"github.com/chaincode_learning_go/chaincode/token/domain/resp"
	"github.com/chaincode_learning_go/chaincode/token/repo/iterf"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type TokenLogService struct {
	TokenLogRepository iterf.ITokenLogRepository
}

func (service *TokenLogService) PaginateTokenLogByFromUserName(stub shim.ChaincodeStubInterface,
	request *protos.PaginateTokenLogByFromUserNameRequest) *protos.PaginateTokenLogByFromUserNameResponse {
	query := []string{request.FromUserName}
	pageSize := request.PageSize
	bookMark := request.BookMark

	results, newBookMark, err := service.TokenLogRepository.PaginateTokenLogByFromUserName(stub, query, pageSize, bookMark)

	if err != nil {
		return &protos.PaginateTokenLogByFromUserNameResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	finalResults := make([]*protos.TokenLogPB, 0)
	for _, result := range results {
		finalResults = append(finalResults, result.ToProtoBufObj())
	}
	return &protos.PaginateTokenLogByFromUserNameResponse{
		CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS),
		PageSize:       pageSize,
		Results:        finalResults,
		BookMark:       newBookMark,
	}
}

func (service *TokenLogService) PaginateTokenLogByToUserName(stub shim.ChaincodeStubInterface,
	request *protos.PaginateTokenLogByToUserNameRequest) *protos.PaginateTokenLogByToUserNameResponse {
	query := []string{request.ToUserName}
	pageSize := request.PageSize
	bookMark := request.BookMark

	results, newBookMark, err := service.TokenLogRepository.PaginateTokenLogByToUserName(stub, query, pageSize, bookMark)

	if err != nil {
		return &protos.PaginateTokenLogByToUserNameResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	finalResults := make([]*protos.TokenLogPB, 0)
	for _, result := range results {
		finalResults = append(finalResults, result.ToProtoBufObj())
	}
	return &protos.PaginateTokenLogByToUserNameResponse{
		CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS),
		PageSize:       pageSize,
		Results:        finalResults,
		BookMark:       newBookMark,
	}
}
