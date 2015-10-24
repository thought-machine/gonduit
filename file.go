package gonduit

import (
	"github.com/etcinit/gonduit/requests"
	"github.com/etcinit/gonduit/responses"
)

// FileDownload performs a call to file.download.
func (c *Conn) FileDownload(req requests.FileDownloadRequest) (*responses.FileDownloadResponse, error) {
	var res responses.FileDownloadResponse

	req.Session = c.Session

	if err := c.Call("file.download", &req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
