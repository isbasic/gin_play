package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	datab "github.com/isbasic/gin_play/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func addPicRoutes(rg *gin.RouterGroup) {
	pic := rg.Group("/pic")

	pic.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pic")
	})

	pic.GET("/uploadfile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploadImg.tmpl", gin.H{
			"title": "图片上传",
		})
	})

	pic.GET("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	pic.GET("/list", func(c *gin.Context) {
		var p datab.Pg
		dsn := p.DSN()
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		var bint []datab.BIN_TEST
		tx := db.Table("bin_tests").Scan(&bint)
		summary := fmt.Sprintf("db: %s,rows: %d", tx.Name(), tx.RowsAffected)

		c.HTML(http.StatusOK, "imgList.tmpl", gin.H{
			"title":   "图片列表",
			"content": bint,
			"tx":      summary,
		})
	})

	pic.GET("/getpic/:id", func(c *gin.Context) {
		var p datab.Pg
		dsn := p.DSN()
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		id := c.Params.ByName("id")
		var bint datab.BIN_TEST
		db.First(&bint, id).Scan(&bint)

		c.Writer.WriteString(string(bint.BData))
	})

	pic.GET("/getpic", func(c *gin.Context) {
		c.JSON(http.StatusOK, "just for fun")
	})
}
