package domains

import (
	i "github.com/OriishiTakahiro/HINAMe-back/infrastructures"
	"log"
)

type Shelter struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	State     string  `db:"state"`
}

func NewShelter(name string, latitude float64, longitude float64) (shelter Shelter) {
	return Shelter{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
		State:     "safe",
	}
}

func (shelter Shelter) CreateTableIfNotExists() error {
	sql := `
		CREATE TABLE IF NOT EXISTS hiname.shelters (
			id 			INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name 		VARCHAR(256) NOT NULL,
			latitude	FLOAT NOT NULL,
			longitude	FLOAT NOT NULL,
			state		ENUM('safe', 'filled', 'danger', 'unavailable') NOT NULL DEFAULT 'safe'
		) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
	`
	if _, err := i.DB.Exec(sql); err != nil {
		return err
	}
	return nil
}

func SheltersInRectangle(minLat float64, minLon float64, maxLat float64, maxLon float64) (shelters []Shelter, e error) {
	sql := `
		SELECT id, name, latitude, longitude, state FROM shelters WHERE 
		(latitude BETWEEN ? AND ?) AND (longitude BETWEEN ? AND ?);
		`
	e = i.DB.Select(&shelters, sql, minLat, maxLat, minLon, maxLon)
	if shelters == nil {
		shelters = []Shelter{}
	}
	return
}

func SheltersInIDRange(minID int, maxID int) (shelters []Shelter, e error) {
	sql := `SELECT id, name, latitude, longitude, state FROM shelters WHERE id BETWEEN ? and ?;`
	e = i.DB.Select(&shelters, sql, minID, maxID)
	log.Println(minID < maxID)
	log.Println(sql)
	if shelters == nil {
		shelters = []Shelter{}
	}
	return
}
