package goalMakeHandler

import (
	"log"
	"net/http"
	"regexp"
)

// HTTP request handler
func HandleRequest(function func(http.ResponseWriter, *http.Request), stringValidPath string) http.HandlerFunc {
	// Sanitize URL path using regexp
	var validPath = regexp.MustCompile(stringValidPath)
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		// HTTP request webroot then open home page
		if request.URL.Path == "/" {
			function(responseWriter, request)
			return
		}
		// HTTP request not webroot then sanitize URL path
		cleanSlicedUri := validPath.FindStringSubmatch(request.URL.Path)
		// URL path not clean
		if cleanSlicedUri == nil {
			// Give 404 response code
			http.Error(responseWriter, request.URL.Path + " is not found", http.StatusNotFound)
			// Log ip address for further process
			log.Println(
				"[error: HandleRequest()] not valid url path [",
				request.URL.Path,
				"] requested from",
				request.RemoteAddr,
			)
			return
		}
		// If URL path is clean then open page handler
		function(responseWriter, request)
	}
}

