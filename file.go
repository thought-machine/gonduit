package gonduit

import (
	"github.com/thought-machine/gonduit/requests"
	"github.com/thought-machine/gonduit/responses"
)

// FileDownload performs a call to file.download.
func (c *Conn) FileDownload(
	req requests.FileDownloadRequest,
) (*responses.FileDownloadResponse, error) {
	var res responses.FileDownloadResponse

	if err := c.Call("file.download", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// FileSearch performs a call to file.search
func (c *Conn) FileSearch(req requests.SearchRequest) (*responses.SearchResponse, error) {
	var res responses.SearchResponse
	if err := c.Call("file.search", &req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
