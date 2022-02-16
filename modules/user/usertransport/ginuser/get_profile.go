package ginuser

import (
	"FirstProject/common"
	"FirstProject/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(ctx component.AppContext) func(ctx2 *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
