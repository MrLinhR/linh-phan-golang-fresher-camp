package main

import (
	"FirstProject/component"
	"FirstProject/modules/note/notemodel"
	"FirstProject/modules/note/notetransport/ginnote"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	test := notemodel.Note{
		Title:   "Test",
		Content: "Demo",
	}

	jsondata, err := json.Marshal(test)
	fmt.Println(string(jsondata))
	DBConnectionStr := "root:sa123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DBConnectionStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(db, err)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	appCtx := component.NewAppContext(db)

	notes := r.Group("/notes")
	{
		notes.POST("", ginnote.CreateNote(appCtx))

		notes.GET("/:id", ginnote.ListNoteById(appCtx))

		notes.GET("", ginnote.ListNote(appCtx))

		notes.PATCH("/:id", ginnote.UpdateNoteById(appCtx))

		notes.DELETE("/:id", ginnote.DeleteNoteById(appCtx))
		//	notes.GET("/:id", func(c *gin.Context) {
		//		id, err := strconv.Atoi(c.Param("id"))
		//
		//		if err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//
		//		var data Note
		//		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//		c.JSON(http.StatusOK, data)
		//	})
		//	notes.PATCH("/:id", func(c *gin.Context) {
		//		id, err := strconv.Atoi(c.Param("id"))
		//
		//		if err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//
		//		var data Note
		//
		//		if err := c.ShouldBind(&data); err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//
		//		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//
		//		c.JSON(http.StatusOK, data)
		//	})
		//
		//	notes.DELETE("/:id", func(c *gin.Context) {
		//		id, err := strconv.Atoi(c.Param("id"))
		//
		//		if err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//
		//		if err := db.Table(Note{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		//			c.JSON(401, map[string]interface{}{
		//				"error": err.Error(),
		//			})
		//			return
		//		}
		//
		//		c.JSON(http.StatusOK, gin.H{"ok": 1})
		//	})
	}
	r.Run()
}
