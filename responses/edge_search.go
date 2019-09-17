package responses

//EdgeSearchResponse is the result of calling edge.search and is a list of object relationships.
type EdgeSearchResponse struct {
	Data   []Edge `json:"data"`
	Cursor struct {
		Limit  int         `json:"limit"`
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
	} `json:"cursor"`
}

type Edge struct {
	SourcePhid      string `json:"sourcePHID"`
	DestinationPhid string `json:"destinationPHID"`
	EdgeType        string `json:"edgeType"`
}
