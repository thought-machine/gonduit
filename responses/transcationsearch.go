package responses

import "github.com/thought-machine/gonduit/util"

type TransactionSearchResponse struct {
	Data   []Data `json:"data"`
	Cursor struct {
		Limit  int         `json:"limit"`
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
	} `json:"cursor"`
}

type Data struct {
	ID           int        `json:"id"`
	Phid         string     `json:"phid"`
	Type         string     `json:"type"`
	AuthorPHID   string     `json:"authorPHID"`
	ObjectPHID   string     `json:"objectPHID"`
	DateCreated  util.UnixTimestamp        `json:"dateCreated"`
	DateModified util.UnixTimestamp        `json:"dateModified"`
	GroupID      string     `json:"groupID"`
	Comments     []Comments `json:"comments"`
	Fields       map[string]interface{}     `json:"fields,omitempty"`
}

type Comments struct {
	ID           int     `json:"id"`
	Phid         string  `json:"phid"`
	Version      int     `json:"version"`
	AuthorPHID   string  `json:"authorPHID"`
	DateCreated  util.UnixTimestamp     `json:"dateCreated"`
	DateModified util.UnixTimestamp     `json:"dateModified"`
	Removed      bool    `json:"removed"`
	Content      map[string]interface{} `json:"content"`
}