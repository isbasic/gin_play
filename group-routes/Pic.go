package routes

import (
	"fmt"
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

	pic.GET("/upload", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pic commit")
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

	pic.GET("/:id", func(c *gin.Context) {
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
