package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamnk74/todolist-mysql-go/dto"
)

func HandleApiError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.ByType(gin.ErrorTypeAny)

		if len(errors) > 0 {
			log.Println("Error ", errors)
			var apiErrors []*dto.ApiError
			for _, errObj := range errors {
				err := errObj.Err
				var apiError *dto.ApiError
				switch err.(type) {
				case *dto.ApiError:
					apiError = err.(*dto.ApiError)
				default:
					apiError = &dto.ApiError{
						Status: http.StatusBadRequest,
						Code:   "ERR-400",
						Title:  http.StatusText(http.StatusBadRequest),
						Detail: err.Error(),
					}
				}
				apiErrors = append(apiErrors, apiError)
			}

			c.IndentedJSON(int(apiErrors[0].Status), gin.H{
				"errors": apiErrors,
			})
			c.Abort()
			return
		}

		// c.AbortWithStatusJSON(500, []dto.ApiError{
		// 	{
		// 		Code:   "ERR-500",
		// 		Status: http.StatusInternalServerError,
		// 		Title:  http.StatusText(http.StatusInternalServerError),
		// 		Detail: http.StatusText(http.StatusInternalServerError),
		// 	},
		// })
		return
	}
}
