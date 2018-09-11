package domain

import "time"

type Board struct {
	ID        int32     `db:"id"`
	ShelterID int32     `db:"shelter_id"`
	HTML      string    `db:"html"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewBoard(shelterID int32) Board {
	return Board{ShelterID: shelterID, HTML: "", UpdatedAt: time.Now()}
}

func (board Board) CreateTableIfNotExists() string {
	return `
		CREATE TABLE IF NOT EXISTS hiname.boards (
			id 			INT NOT NULL PRIMARY KEY KEY AUTO_INCREMENT,
			shelter_id 	INT NOT NULL UNIQUE,
			html		TEXT NOT NULL DEFAULT '<h3>ようこそ</h3>',
			updated_at	TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
			FOREIGN KEY(shelter_id)	REFERENCES hiname.shelters(id)
		) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
	`
}
