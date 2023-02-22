package nestgo

import "net/http"

type Context struct {
	r       *http.Request
	w       http.ResponseWriter
	headers map[string]string
	body    interface{}
}
