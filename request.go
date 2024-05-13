package util

import (
	"net/http"
	"strings"
	
	"github.com/creamsensation/util/constant/contentType"
	"github.com/creamsensation/util/constant/header"
)

func IsRequestMultipart(req *http.Request) bool {
	return strings.Contains(req.Header.Get(header.ContentType), contentType.MultipartForm)
}
