package ginnote

import (
	"FirstProject/common"
	"FirstProject/component"
	"FirstProject/modules/note/notebiz"
	"FirstProject/modules/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteNoteById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewDeleteNoteBiz(store)
		id, ty := strconv.Atoi(c.Param("id"))
		if ty != nil {
			c.JSON(401, map[string]interface{}{
				"error": ty.Error(),
			})
			return
		}
		result, err := biz.DeleteNote(c.Request.Context(), id)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
