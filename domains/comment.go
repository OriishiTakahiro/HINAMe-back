package domain

import "time"

type Comment struct {
	ID        int32     `id:"id"`
	BoardID   int32     `db:"board_id"`
	ParentID  int32     `db:"parent_id"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	Body      string    `db:"body"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewComment(boardID int32, parentID int32, title string, author string, body string) Comment {
	return Comment{
		BoardID:  boardID,
		ParentID: parentID,
		Title:    title,
		Author:   author,
		Body:     body,
	}
}

func (comment Comment) CreateTableIfNotExists() string {
	return `
		CREATE TABLE IF NOT EXISTS hiname.comments (
			id 			INT NOT NULL PRIMARY KEY KEY AUTO_INCREMENT,
			board_id 	INT NOT NULL,
			parent_id 	INT DEFAULT NULL,
			title		VARCHAR(256) NOT NULL DEFAULT 'タイトルなし',
			author		VARCHAR(256) NOT NULL DEFAULT '住民',
			body		TEXT NOT NULL,
			updated_at	TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
			FOREIGN KEY(board_id)	REFERENCES hiname.boards(id)
		) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
	`
}
