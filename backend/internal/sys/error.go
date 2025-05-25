package sys

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/merynayr/mall-tenants/internal/logger"
	"github.com/merynayr/mall-tenants/internal/sys/codes"
)

// ErrorResponse Структура для обработчки http ошибок
type ErrorResponse struct {
	msg  string
	code codes.Code
	err  error
}

// NewCommonError создаёт новвую ошибку
func NewCommonError(msg string, code codes.Code) *ErrorResponse {
	return &ErrorResponse{msg, code, nil}
}

// Wrap оборачивает оригинальную ошибку в кастомную, сохраняя трассировку
func Wrap(err error, base *ErrorResponse) *ErrorResponse {
	return &ErrorResponse{
		msg:  base.msg,
		code: base.code,
		err:  errors.WithStack(err),
	}
}

// Error возврашает сообщение ошибки
func (r *ErrorResponse) Error() string {
	return r.msg
}

// Code возвращает код ошибки
func (r *ErrorResponse) Code() codes.Code {
	return r.code
}

// Unwrap разбирает ошибку
func (r *ErrorResponse) Unwrap() error {
	return r.err
}

// IsCommonError проверяет на соответствие ошибке
func IsCommonError(err error) bool {
	var ce *ErrorResponse
	return errors.As(err, &ce)
}

// GetCommonError получает ошбику
func GetCommonError(err error) *ErrorResponse {
	var ce *ErrorResponse
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}

// HandleError обрабатывает ошибки и отправляет корректный HTTP-ответ
func HandleError(c *gin.Context, err error) {
	var ce *ErrorResponse
	if errors.As(err, &ce) {
		if ce.err != nil {
			logger.Error(errors.WithStack(ce.err).Error())
		} else {
			logger.Debug(ce.Error())
		}

		c.JSON(int(ce.Code()), gin.H{
			"error": ce.Error(),
			"code":  ce.Code(),
		})
		return
	}

	logger.Error(errors.WithStack(err).Error())
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
		"code":  http.StatusInternalServerError,
	})
}
