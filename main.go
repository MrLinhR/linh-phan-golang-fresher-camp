package main

import (
	"FirstProject/component"
	"FirstProject/component/uploadprovider"
	"FirstProject/middleware"
	"FirstProject/modules/upload/uploadtransport/ginupload"
	"FirstProject/modules/user/usertransport/ginuser"
	"os"

	"FirstProject/modules/note/notetransport/ginnote"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//DBConnectionStr := "root:sa123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(os.Getenv("DBConnectionStr")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()

	s3BucketName := os.Getenv("s3BucketName")
	s3Region := os.Getenv("s3Region")
	s3APIKey := os.Getenv("s3APIKey")
	s3SecretKey := os.Getenv("s3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := component.NewAppContext(db, s3Provider, secretKey)

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))

	v1.POST("/login", ginuser.Login(appCtx))

	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	notes := v1.Group("/notes")
	{
		notes.POST("", ginnote.CreateNote(appCtx))

		notes.GET("/:id", ginnote.ListNoteById(appCtx))

		notes.GET("", ginnote.ListNote(appCtx))

		notes.PATCH("/:id", ginnote.UpdateNoteById(appCtx))

		notes.DELETE("/:id", ginnote.DeleteNoteById(appCtx))
	}
	r.Run()
}
