package ee

import (
	"io"
	"net/http"
)

func CloseHTTPResponse(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
}
