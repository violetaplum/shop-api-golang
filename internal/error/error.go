package shop_error

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// HTTP 에러 응답 구조체
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// gRPC 에러를 HTTP 에러로 변환하는 함수
func HandleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		// gRPC 상태가 아닌 경우 기본 에러 반환
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		})
		return
	}

	// gRPC 코드를 HTTP 상태코드로 변환
	var httpStatus int
	var errorCode string

	switch st.Code() {
	case codes.NotFound:
		httpStatus = http.StatusNotFound
		errorCode = "NOT_FOUND"
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
		errorCode = "INVALID_ARGUMENT"
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
		errorCode = "ALREADY_EXISTS"
	case codes.PermissionDenied:
		httpStatus = http.StatusForbidden
		errorCode = "PERMISSION_DENIED"
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
		errorCode = "UNAUTHENTICATED"
	case codes.ResourceExhausted:
		httpStatus = http.StatusTooManyRequests
		errorCode = "RESOURCE_EXHAUSTED"
	case codes.FailedPrecondition:
		httpStatus = http.StatusPreconditionFailed
		errorCode = "FAILED_PRECONDITION"
	case codes.Unavailable:
		httpStatus = http.StatusServiceUnavailable
		errorCode = "SERVICE_UNAVAILABLE"
	default:
		httpStatus = http.StatusInternalServerError
		errorCode = "INTERNAL_ERROR"
	}

	c.JSON(httpStatus, ErrorResponse{
		Code:    errorCode,
		Message: st.Message(),
	})
}
