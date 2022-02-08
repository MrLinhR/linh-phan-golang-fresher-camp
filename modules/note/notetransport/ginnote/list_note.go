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

func ListNote(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewListNoteBiz(store)
		result, err := biz.ListNote(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

func ListNoteById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := notestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := notebiz.NewListNoteBiz(store)
		id, ty := strconv.Atoi(c.Param("id"))
		if ty != nil {
			c.JSON(401, map[string]interface{}{
				"error": ty.Error(),
			})
			return
		}
		result, err := biz.ListNoteById(c.Request.Context(), id, &paging)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
