package application

import (
	"fmt"
	d "github.com/OriishiTakahiro/HINAMe-back/domains"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	sc "strconv"
)

// GetSheltersIn get shelters in given id range
func GetSheltersIn(c *gin.Context) {
	minID, err2 := sc.Atoi(c.Param("min_id"))
	maxID, err1 := sc.Atoi(c.Param("max_id"))

	if err1 != nil || err2 != nil {
		msg := fmt.Sprintf("Parameters has invalid type (%s, %s)", c.Param("min_id"), c.Param("max_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}

	shelters, err := d.SheltersInIDRange(minID, maxID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Response is failed")
		return
	}
	c.JSON(http.StatusOK, shelters)
	return
}

// GetSheltersInRect get shelters in given rect
func GetSheltersInRect(c *gin.Context) {
	minLat, err3 := sc.ParseFloat(c.Param("min_latitude"), 64)
	minLon, err4 := sc.ParseFloat(c.Param("min_longitude"), 64)
	maxLat, err1 := sc.ParseFloat(c.Param("max_latitude"), 64)
	maxLon, err2 := sc.ParseFloat(c.Param("max_longitude"), 64)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		msg := fmt.Sprintf(
			"Parameters has invalid type (%s, %s, %s, %s)",
			c.Param("max_latitude"),
			c.Param("max_longitude"),
			c.Param("min_latitude"),
			c.Param("min_longitude"),
		)
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}
	shelters, err := d.SheltersInRectangle(minLat, minLon, maxLat, maxLon)
	if err != nil {
		c.String(http.StatusInternalServerError, "Request is failed")
		return
	}
	c.JSON(http.StatusOK, shelters)
	return
}

// GetShelterNamesIn get shelter name list in given id range
func GetShelterNamesIn(c *gin.Context) {
	minID, err2 := sc.Atoi(c.Param("min_id"))
	maxID, err1 := sc.Atoi(c.Param("max_id"))
	if err1 != nil || err2 != nil {
		msg := fmt.Sprintf("Parameters has invalid type (%s, %s)", c.Param("min_id"), c.Param("max_id"))
		log.Println(msg)
		c.String(http.StatusBadRequest, msg)
		return
	}
	shelters, err := d.SheltersInIDRange(minID, maxID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, "Request is failed")
		return
	}
	c.JSON(http.StatusOK, shelters)
}
