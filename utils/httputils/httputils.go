package httputils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIntQueryParamValue - Get the value of a query param as an int
func GetIntQueryParamValue(c *gin.Context, key string) (int, error) {
	return strconv.Atoi(c.Param(key))
}

// GetIntPostValue - Get the value of a request post as an int
func GetIntPostValue(c *gin.Context, key string) (int, error) {
	return strconv.Atoi(c.PostForm(key))
}

// GetDefaultIntPostValue - Get the value of a request post as an int that fallsback to a default
func GetDefaultIntPostValue(c *gin.Context, key string, defaultValue int) (int, error) {
	return strconv.Atoi(c.DefaultPostForm(key, strconv.Itoa(defaultValue)))
}

// GetFloatPostValue - Get the value of a request post as a float64
func GetFloatPostValue(c *gin.Context, key string) (float64, error) {
	return strconv.ParseFloat(c.PostForm(key), 64)
}

// GetDefaultFloatPostValue - Get the value of a request post as a float64 that fallsback to a default
func GetDefaultFloatPostValue(c *gin.Context, key string, defaultValue float64) (float64, error) {
	return strconv.ParseFloat(c.DefaultPostForm(key, strconv.FormatFloat(defaultValue, 'e', 2, 64)), 64)
}

// HandleErrorOrSuccessResponse - Return either a error or success response
func HandleErrorOrSuccessResponse(c *gin.Context, e error, data interface{}, onBeforeResponse func()) {
	if onBeforeResponse == nil {
		onBeforeResponse = func() {}
	}

	if e != nil {
		fmt.Println(e.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   e.Error(),
		})
	} else {
		onBeforeResponse()
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    data,
		})
	}
}