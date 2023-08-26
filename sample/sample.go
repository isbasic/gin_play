package sample

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	model "github.com/isbasic/gin_play/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = make(map[string]string)

func B64Encode(in []byte) string {
	res := base64.StdEncoding.EncodeToString(in)
	return res
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"B64": B64Encode,
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/zhan", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "老詹快乐水",
			"content": "经典款促销，一件3000.",
		})
	})

	r.GET("/zhan/:slogan", func(c *gin.Context) {
		slogan := c.Params.ByName("slogan")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "老詹快乐水",
			"content": slogan,
		})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/goscreen", func(c *gin.Context) {
		fp := "static/Go-ScreenShot.jpeg"
		f, _ := ioutil.ReadFile(fp)
		imgContent := string(f)
		c.Writer.WriteString(imgContent)
	})

	r.GET("/goimg", func(c *gin.Context) {
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		var bint []model.BIN_TEST
		tx := db.Table("bin_tests").Scan(&bint)
		summary := fmt.Sprintf("db: %s,rows: %d", tx.Name(), tx.RowsAffected)

		c.HTML(http.StatusOK, "imgView.tmpl", gin.H{
			"title":   "图片查看",
			"content": bint,
			"tx":      summary,
		})
	})

	r.GET("/goimg/:id", func(c *gin.Context) {
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		id := c.Params.ByName("id")

		var bint []model.BIN_TEST
		tx := db.First(&bint, id).Scan(&bint)
		summary := fmt.Sprintf("db: %s,rows: %d", tx.Name(), tx.RowsAffected)

		c.HTML(http.StatusOK, "imgView.tmpl", gin.H{
			"title":   "图片查看",
			"content": bint,
			"tx":      summary,
		})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// 将输出：while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	r.GET("/savepic", func(c *gin.Context) {
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		name := db.Name()
		fp := "static/Go-ScreenShot.jpeg"
		f, err := ioutil.ReadFile(fp)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"db": fmt.Sprintf("%s read failed, err: %s.", name, err)})
			return
		}

		gosc := model.BIN_TEST{BData: f}
		res := db.Create(&gosc)
		if res.Error != nil {
			c.JSON(http.StatusOK, gin.H{"db": fmt.Sprintf("%s save to database failed, err: %s.", name, res.Error)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"db": fmt.Sprintf("%s opend, write %d rows,id: %d.", name, res.RowsAffected, gosc.Sampleid)})
	})

	r.GET("/uploadpic/:name", func(c *gin.Context) {
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		name := c.Params.ByName("name")
		fp := fmt.Sprintf("static/%s", name)
		f, err := ioutil.ReadFile(fp)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"db": fmt.Sprintf("%s read failed, err: %s.", name, err)})
			return
		}

		gosc := model.BIN_TEST{BData: f}
		res := db.Create(&gosc)
		if res.Error != nil {
			c.JSON(http.StatusOK, gin.H{"db": fmt.Sprintf("%s save to database failed, err: %s.", name, res.Error)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"db": fmt.Sprintf("%s opend, write %d rows,id: %d.", name, res.RowsAffected, gosc.Sampleid)})
	})

	r.GET("/getpic/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		var bint model.BIN_TEST
		tx := db.First(&bint, id).Scan(&bint)
		// c.Writer.WriteString(string(gosc.BData))
		c.JSON(http.StatusOK, gin.H{"result count": tx.RowsAffected, "data": bint})
	})

	r.GET("/viewpic", func(c *gin.Context) {
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		var bint model.BIN_TEST
		db.First(&bint, 1).Scan(&bint)
		// c.Writer.WriteString(string(gosc.BData))
		c.Writer.WriteString(string(bint.BData))
	})

	r.GET("/viewpic/:id", func(c *gin.Context) {
		dsn := "host=localhost user=mac password=Yc_19860717 dbname=mac port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		id := c.Params.ByName("id")
		var bint model.BIN_TEST
		db.First(&bint, id).Scan(&bint)
		// c.Writer.WriteString(string(gosc.BData))
		c.Writer.WriteString(string(bint.BData))
	})

	r.GET("/fssy", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"nick": "肥仔快乐水"})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	r.POST("/upload", func(c *gin.Context) {

	})

	r.GET("/dict/:catalog", func(c *gin.Context) {

	})

	r.GET("/upload/newfile", func(c *gin.Context) {

	})

	return r
}
