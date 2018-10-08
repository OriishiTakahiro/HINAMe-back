package domains

import (
	i "github.com/OriishiTakahiro/HINAMe-back/infrastructures"
)

type Comment struct {
	ID        int    `id:"id"`
	BoardID   int    `db:"shelter_id"`
	ParentID  int    `db:"parent_id"`
	Title     string `db:"title"`
	Author    string `db:"author"`
	Body      string `db:"body"`
	UpdatedAt int    `db:"updated_at"`
}

// NewComment create instantiate Comment
func NewComment(shelterID int, parentID int, title string, author string, body string) Comment {
	return Comment{
		BoardID:  shelterID,
		ParentID: parentID,
		Title:    title,
		Author:   author,
		Body:     body,
	}
}

// CreateTableIfNotExists create comments table
func (comment Comment) CreateTableIfNotExists() error {
	sql := `
		CREATE TABLE IF NOT EXISTS hiname.comments (
			id 			INT NOT NULL PRIMARY KEY KEY AUTO_INCREMENT,
			shelter_id 	INT NOT NULL,
			parent_id 	INT DEFAULT -1,
			title		VARCHAR(256) NOT NULL DEFAULT 'タイトルなし',
			author		VARCHAR(256) NOT NULL DEFAULT '住民',
			body		TEXT NOT NULL,
			updated_at	INT NOT NULL DEFAULT UNIX_TIMESTAMP(NOW()),
			FOREIGN KEY(shelter_id) REFERENCES hiname.shelters(id)
		) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
	`
	_, e := i.DB.Exec(sql)
	return e
}

func (comment Comment) GetByShelterID(shelterID int) (comments []Comment, e error) {
	sql := `
		SELECT 
		  id, parent_id, title, author, body, updated_at 
		FROM 
		  hiname.comments 
		WHERE 
		  shelter_id = ? 
		ORDER BY updated_at DESC;
	`
	e = i.DB.Select(&comments, sql, shelterID)
	if comments == nil {
		comments = []Comment{}
	}
	return
}

// GetReplies get replies by pareint id
func (comment Comment) GetReplies(parentID int) (comments []Comment, e error) {
	sql := `SELECT id, title, author, body, updated_at FROM hiname.comments WHERE parent_id = ?;`
	e = i.DB.Select(&comments, sql, parentID)
	if comments == nil {
		comments = []Comment{}
	}
	return
}

// Upload new comment
func (comment Comment) Upload(shelterID, parentID int, title, author, body string) (e error) {
	if parentID < 0 {
		sql := `INSERT INTO hiname.comments (shelter_id, title, author, body) VALUES (?, ?, ?, ?)`
		_, e = i.DB.Exec(sql, shelterID, title, author, body)
	} else {
		sql := `INSERT INTO hiname.comments (shelter_id, parent_id, title, author, body) VALUES (?, ?, ?, ?, ?)`
		_, e = i.DB.Exec(sql, shelterID, parentID, title, author, body)
	}
	return
}

// Update update a comment
func (comment Comment) Update() (e error) {
	sql := `
	UPDATE hiname.comments SET (shelter_id = ?, parent_id = ?, title = ?, author = ?, body = ?, updated_at = UNIX_TIMESTAMP(NOW())) where id = ?;
	`
	_, e = i.DB.Exec(sql, comment.BoardID, comment.ParentID, comment.Title, comment.Author, comment.Body, comment.ID)
	return
}
