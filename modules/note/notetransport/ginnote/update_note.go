package ginnote

import (
	"FirstProject/common"
	"FirstProject/component"
	"FirstProject/modules/note/notebiz"
	"FirstProject/modules/note/notemodel"
	"FirstProject/modules/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateNoteById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data notemodel.Note
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewUpdateNoteBiz(store)
		id, ty := strconv.Atoi(c.Param("id"))
		if ty != nil {
			c.JSON(401, map[string]interface{}{
				"error": ty.Error(),
			})
			return
		}
		err := biz.UpdateNote(c.Request.Context(), &data, id)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
