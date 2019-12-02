package impl

import (
	"github.com/chaincode_learning_go/chaincode/token/consts"
	"github.com/chaincode_learning_go/chaincode/token/domain/protos"
	"github.com/chaincode_learning_go/chaincode/token/domain/resp"
	"github.com/chaincode_learning_go/chaincode/token/repo/iterf"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type TokenService struct {
	TokenCreationRepository iterf.ITokenCreationRepository
	TokenLogRepository      iterf.ITokenLogRepository
	TokenBalanceRepository  iterf.ITokenBalanceRepository
}

func (service *TokenService) Ping(stub shim.ChaincodeStubInterface, request *protos.PingRequest) *protos.PingResponse {
	return &protos.PingResponse{CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS)}
}

func (service *TokenService) CreateToken(stub shim.ChaincodeStubInterface, request *protos.CreateTokenRequest) *protos.CreateTokenResponse {

	err := service.TokenCreationRepository.CreateToken(stub, request.TokenName, request.MaxAmount, request.Creator, request.Issuer)

	if err != nil {
		return &protos.CreateTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	return &protos.CreateTokenResponse{CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS)}
}

func (service *TokenService) IssueToken(stub shim.ChaincodeStubInterface, request *protos.IssueTokenRequest) *protos.IssueTokenResponse {
	requestId := request.CommonRequest.RequestId
	userName := request.UserName
	tokenName := request.TokenName
	tokenAmount := request.TokenAmount
	err := service.TokenCreationRepository.UpdateTokenIssueAmount(stub, userName, tokenName, tokenAmount)

	if err != nil {
		return &protos.IssueTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	err = service.TokenBalanceRepository.AddBalance(stub, userName, tokenName, tokenAmount)

	if err != nil {
		return &protos.IssueTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	err = service.TokenLogRepository.CreateTokenLog(stub, requestId, userName, tokenName, tokenAmount)

	if err != nil {
		return &protos.IssueTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	return &protos.IssueTokenResponse{CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS)}
}

func (service *TokenService) TransferToken(stub shim.ChaincodeStubInterface, request *protos.TransferTokenRequest) *protos.TransferTokenResponse {
	requestId := request.CommonRequest.RequestId
	fromUserName := request.FromUserName
	toUserName := request.ToUserName
	tokenName := request.TokenName
	tokenAmount := request.TokenAmount

	err := service.TokenBalanceRepository.DeductBalance(stub, fromUserName, tokenName, tokenAmount)

	if err != nil {
		return &protos.TransferTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	err = service.TokenBalanceRepository.AddBalance(stub, toUserName, tokenName, tokenAmount)

	if err != nil {
		return &protos.TransferTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	err = service.TokenLogRepository.CreateTokenLog(stub, requestId, fromUserName, tokenName, -tokenAmount)

	if err != nil {
		return &protos.TransferTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	err = service.TokenLogRepository.CreateTokenLog(stub, requestId, toUserName, tokenName, tokenAmount)

	if err != nil {
		return &protos.TransferTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	return &protos.TransferTokenResponse{CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS)}
}

func (service *TokenService) GetToken(stub shim.ChaincodeStubInterface, request *protos.GetTokenRequest) *protos.GetTokenResponse {

	userBalance, err := service.TokenBalanceRepository.GetBalance(stub, request.UserName, request.TokenName)

	if err != nil {
		return &protos.GetTokenResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	return &protos.GetTokenResponse{
		CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS),
		Result:         userBalance,
	}
}

func (service *TokenService) PaginateTokenByUserName(stub shim.ChaincodeStubInterface,
	request *protos.PaginateTokenByUserNameRequest) *protos.PaginateTokenByUserNameResponse {
	query := []string{request.UserName}
	pageSize := request.PageSize
	bookMark := request.BookMark

	results, newBookMark, err := service.TokenBalanceRepository.PaginateBalanceByUserName(stub, query, pageSize, bookMark)

	if err != nil {
		return &protos.PaginateTokenByUserNameResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	return &protos.PaginateTokenByUserNameResponse{
		CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS),
		PageSize:       pageSize,
		Results:        results,
		BookMark:       newBookMark,
	}
}

func (service *TokenService) PaginateTokenByTokenName(stub shim.ChaincodeStubInterface,
	request *protos.PaginateTokenByTokenNameRequest) *protos.PaginateTokenByTokenNameResponse {
	query := []string{request.TokenName}
	pageSize := request.PageSize
	bookMark := request.BookMark

	results, newBookMark, err := service.TokenBalanceRepository.PaginateBalanceByTokenName(stub, query, pageSize, bookMark)

	if err != nil {
		return &protos.PaginateTokenByTokenNameResponse{
			CommonResponse: resp.BuildErrorCommonResponse(consts.RESPONSE_FAILURE, consts.GENERAL_ERROR_CODE, err.Error()),
		}
	}

	return &protos.PaginateTokenByTokenNameResponse{
		CommonResponse: resp.BuildSuccessCommonResponse(consts.RESPONSE_SUCCESS),
		PageSize:       pageSize,
		Results:        results,
		BookMark:       newBookMark,
	}
}
