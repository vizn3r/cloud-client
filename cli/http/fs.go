package http

import (
	"net/http"
)

// This module will also be used by the desktop daemon in the future, so try to make it modular

// Now the fun begins :P
// Server requires:
//
// POST /file/init to init the uploading session
// Then send required chunks of file to POST /file/:sid/:chunk
// The server will respond with the required chunk ID
// When all chunks are sent, the server will respond with chunk ID -1

func UploadFile(filePath string, fileName string) string {
	res, err := http.Post(cfg.ServerHost, "text/json", nil)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return ""
	}
	return ""
}
