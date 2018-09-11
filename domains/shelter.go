package domain

type Shelter struct {
	ID        int32   `db:"id"`
	Name      string  `db:"name"`
	Latitude  float32 `db:"latitude"`
	Longitude float32 `db:"longitude"`
	State     string  `db:"state"`
}

func NewShelter(name string, latitude float32, longitude float32) (shelter Shelter) {
	return Shelter{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
		State:     "safe",
	}
}

func (shelter Shelter) CreateTableIfNotExists() string {
	return `
		CREATE TABLE IF NOT EXISTS hiname.shelters (
			id 			INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
			name 		VARCHAR(256) NOT NULL,
			latitude	FLOAT NOT NULL,
			longitude	FLOAT NOT NULL,
			state		ENUM('safe', 'filled', 'danger', 'unavailable') NOT NULL DEFAULT 'safe'
		) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
	`
}
