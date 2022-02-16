package ginuser

import (
	"FirstProject/common"
	"FirstProject/component"
	"FirstProject/component/hasher"
	"FirstProject/modules/user/userbusiness"
	"FirstProject/modules/user/usermodel"
	"FirstProject/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) func(ctx2 *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		repo := userbusiness.NewRegisterBusiness(store, md5)
		if err := repo.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
