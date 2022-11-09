package routers

import (
	"wms/controllers"

	"github.com/gin-gonic/gin"
)

func CouchDBRouter(r *gin.Engine) {
	r.Static("assets/", "./assets")

	master := r.Group("/couch")
	{
		master.POST("/getdoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.Find(c))
		})
		master.POST("/insertdoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.Insert(c))
		})
		master.POST("/updatedoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.Update(c))
		})
		master.POST("/deletedoc", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.Delete(c))
		})
		master.POST("/bulkdocs", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.BulkDocs(c))
		})
		master.POST("/createdb", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.CreateDB(c))
		})
		master.POST("/deletedb", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.DeleteDB(c))
		})
		master.POST("/createuserdb", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.CreateUserDB(c))
		})

		master.POST("/getview", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.GetView(c))
		})
		master.POST("/registercompany", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.RegisterCompany(c))
		})
		master.POST("/setredis", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.SetRedis(c))
		})
		master.POST("/getredis", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.GetRedis(c))
		})
		master.POST("/checksession", func(c *gin.Context) {
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.String(200, controllers.CheckSession(c))
		})
	}
}
