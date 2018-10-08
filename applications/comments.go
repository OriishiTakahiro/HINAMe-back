package application

import (
	"fmt"
	d "github.com/OriishiTakahiro/HINAMe-back/domains"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	sc "strconv"
)

// GetComments get all comments by shelterID
func GetComments(c *gin.Context) {
	shelterID, err := sc.Atoi(c.Param("shelter_id"))
	if err != nil {
		msg := fmt.Sprintf("Parameters has invalid type %s", c.Param("shelter_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}

	comments, err := d.Comment{}.GetByShelterID(shelterID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Request is failed.")
		return
	}
	c.JSON(http.StatusOK, comments)
	return
}

// GetReplies get all replies by comment id
func GetReplies(c *gin.Context) {
	parentID, err := sc.Atoi(c.Param("parent_id"))
	if err != nil {
		msg := fmt.Sprintf("Parameters has invalid type %s", c.Param("parent_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}

	replies, err := d.Comment{}.GetReplies(parentID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Request is failed")
		return
	}

	c.JSON(http.StatusOK, replies)
	return
}

func UploadComment(c *gin.Context) {
	type Params struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Body   string `json:"body"`
	}

	shelterID, err := sc.Atoi(c.Param("shelter_id"))
	if err != nil {
		msg := fmt.Sprintf("Parameters has invalid type %s", c.Param("shelter_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}
	parentID, err := sc.Atoi(c.Param("parent_id"))
	if err != nil {
		parentID = -1
	}

	var params Params
	if err := c.BindJSON(&params); err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Request is failed")
		return
	}

	err = d.Comment{}.Upload(shelterID, parentID, params.Title, params.Author, params.Body)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Response is failed")
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
	return
}
