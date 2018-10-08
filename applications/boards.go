package application

import (
	"fmt"
	d "github.com/OriishiTakahiro/HINAMe-back/domains"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	sc "strconv"
)

// GetBoardHML get board html by shelter_id
func GetBoardHTML(c *gin.Context) {
	shelterID, err := sc.Atoi(c.Param("shelter_id"))
	if err != nil {
		msg := fmt.Sprintf("Parameters has invalid type %s", c.Param("shelter_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}

	html, err := d.Board{}.GetHTML(shelterID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Response is failed")
		return
	}
	c.String(http.StatusOK, html)
	return
}

func UpdateBoard(c *gin.Context) {
	id, err := sc.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Parameters has invalid type %s", c.Param("shelter_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}

	requestBody := make([]byte, 1024)
	if _, err := c.Request.Body.Read(requestBody); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Response is failed")
		return
	}

	if err := (d.Board{}).Update(id, string(requestBody)); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Response is failed")
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
	return
}
