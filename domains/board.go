package domains

import (
	i "github.com/OriishiTakahiro/HINAMe-back/infrastructures"
	"time"
)

type Board struct {
	ID        int       `db:"id"`
	ShelterID int       `db:"shelter_id"`
	HTML      string    `db:"html"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewBoard(shelterID int) Board {
	return Board{ShelterID: shelterID, HTML: "", UpdatedAt: time.Now()}
}

func (board Board) CreateTableIfNotExists() error {
	sql := `
		CREATE TABLE IF NOT EXISTS hiname.boards (
			id 			INT NOT NULL PRIMARY KEY KEY AUTO_INCREMENT,
			shelter_id 	INT NOT NULL UNIQUE,
			html		TEXT NOT NULL DEFAULT '<h3>ようこそ</h3>',
			updated_at	TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
			FOREIGN KEY(shelter_id)	REFERENCES hiname.shelters(id)
		) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
	`
	_, e := i.DB.Exec(sql)
	return e
}

func (board Board) GetHTML(shelterID int) (html string, e error) {
	sql := ` SELECT (html) FROM hiname.boards WHERE ? = shelter_id;`
	e = i.DB.Get(&html, sql, shelterID)
	return
}

func (board Board) GetIDShelterID(afterFrom time.Time) (boards []Board, e error) {
	sql := ` SELECT (id, shelter_id) FROM hiname.boards (html) WHERE ? > updated_at;`
	e = i.DB.Select(&boards, sql, afterFrom.Format("2006-01-02 15:04:05"))
	if boards == nil {
		boards = []Board{}
	}
	return
}

func (board Board) Update(id int, html string) (e error) {
	sql := `UPDATE hiname.boards FROM html = ? WHERE id = ?;`
	_, e = i.DB.Exec(sql, html, id)
	return
}
