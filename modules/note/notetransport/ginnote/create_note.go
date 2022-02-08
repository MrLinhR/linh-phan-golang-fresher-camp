package ginnote

import (
	"FirstProject/common"
	"FirstProject/component"
	"FirstProject/modules/note/notebiz"
	"FirstProject/modules/note/notemodel"
	"FirstProject/modules/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNote(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data notemodel.NoteCreate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewCreateNoteBiz(store)
		if err := biz.CreateNote(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
