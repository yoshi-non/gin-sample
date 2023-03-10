package main

import (
	"gintut/config"
	"gintut/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()
	//rotuerの宣言
	router := gin.New()
	router.Use(cors.Default())
	//ルータメソッドの宣言 GET,POST,PUT,DELETE("/route_path/:parameter",function)
	router.GET("/get", getAllData)
	router.POST("/post", postData)
	router.PUT("/put", updateData)
	router.DELETE("/delete/:id", deleteData)
	//サーバの起動、IPaddress:portで指定可能
	router.Run("localhost:8080")
}

func getAllData(c *gin.Context) {
	var memo []models.Memo
	//DBから取得 SELECT * FROM memos;
	result := config.DB.Find(&memo)
	//エラーチェック
	if result.Error != nil {
		panic(result.Error)
	}
	c.JSON(
		200,
		gin.H{
			"data": memo,
		})
}

func postData(c *gin.Context) {
	var memo models.Memo
	//送られてきたデータを構造体の形式でバインド
	c.ShouldBindJSON(&memo)
	memo.Updatetime = time.Now()

	//レコードを追加
	result := config.DB.Create(&memo)
	//エラーチェック
	if result.Error != nil {
		panic(result.Error)
	}
	//JSON形式でメッセージ返す
	c.JSON(
		200,
		gin.H{
			"message": "success",
		})
}

func updateData(c *gin.Context) {
	var memo models.Memo
	//データをバインド
	c.ShouldBindJSON(&memo)
	memo.Updatetime = time.Now()
	//DBにセーブ
	config.DB.Save(&memo)
	c.JSON(
		200,
		gin.H{
			"message": "success",
		})
}

func deleteData(c *gin.Context) {
	var memo models.Memo
	//プライマリキーを指定して削除　c.Param() => url パラメタを抽出する
	result := config.DB.Delete(&memo, c.Param("id"))

	//エラーチェック
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(
		200,
		gin.H{
			"message": "success",
		})
}
