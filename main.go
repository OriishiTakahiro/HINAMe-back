package main

import (
	a "github.com/OriishiTakahiro/HINAMe-back/applications"
	d "github.com/OriishiTakahiro/HINAMe-back/domains"
	i "github.com/OriishiTakahiro/HINAMe-back/infrastructures"
	m "github.com/OriishiTakahiro/HINAMe-back/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	if err := i.OpenDB(); err != nil {
		panic(err)
	}
	defer i.CloseDB()
	prepareTables()

	router := gin.Default()
	router.Use(m.CORSMiddleware())

	shelterGroup := router.Group("/api/shelter")
	{
		shelterGroup.GET("/by_id/:min_id/:max_id", a.GetSheltersIn)
		shelterGroup.GET("/by_rect/:min_latitude/:min_longitude/:max_latitude/:max_longitude", a.GetSheltersInRect)
		shelterGroup.GET("/name/by_id/:min_id/:max_id", a.GetShelterNamesIn)
	}
	boardGroup := router.Group("/api/board")
	{
		boardGroup.GET("/html/:shelter_id", a.GetBoardHTML)
		// boardGroup.GET("/shelter_id/:updated_after", a.GetBoardsRecentlyUpdated)
		boardGroup.PUT("/:id", a.UpdateBoard)
	}
	commentGroup := router.Group("/api/comment")
	{
		commentGroup.GET("/by_shelter_id/:shelter_id", a.GetComments)
		commentGroup.GET("/reply/:parent_id", a.GetReplies)
		commentGroup.POST("/:shelter_id/:parent_id", a.UploadComment)
		//commentGroup.PUT("/:id", a.UpdateComment)
		//commentGroup.DELETE("/:id", a.DeleteComment)
	}

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Health OK")
	})

	port, ok := os.LookupEnv("PORT")
	if ok {
		router.Run(":" + port)
	}

	router.Run()
}

func prepareTables() {
	d.Shelter{}.CreateTableIfNotExists()
	d.Comment{}.CreateTableIfNotExists()
	d.Board{}.CreateTableIfNotExists()
}
