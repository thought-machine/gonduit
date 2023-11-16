package responses

import "github.com/samwestmoreland/gonduit/constants"

// SearchResponse is the response of calling one of the .search endpoints.
type SearchResponse struct {
	Data   []SearchData `json:"data"`
	Cursor struct {
		Limit  int                   `json:"limit"`
		After  string                `json:"after"`
		Before string                `json:"before"`
		Order  constants.SearchOrder `json:"order"`
	} `json:"cursor"`
}

// SearchData is a single response item from a search.
type SearchData struct {
	ID          int                               `json:"id"`
	Type        string                            `json:"type"`
	PHID        string                            `json:"phid"`
	Fields      map[string]interface{}            `json:"fields"`
	Attachments map[string]map[string]interface{} `json:"attachments"`
}
