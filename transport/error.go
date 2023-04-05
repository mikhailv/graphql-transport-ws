package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

type gqlResponse struct {
	Errors     gqlerror.List          `json:"errors,omitempty"`
	Data       json.RawMessage        `json:"data"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// SendError sends a best effort error to a raw response writer. It assumes the client can understand the standard
// json error response
func SendError(w http.ResponseWriter, code int, errors ...*gqlerror.Error) {
	w.WriteHeader(code)
	b, err := jsonEncode(&gqlResponse{Errors: errors})
	if err != nil {
		panic(err)
	}
	_, _ = w.Write(b)
}

// SendErrorf wraps SendError to add formatted messages
func SendErrorf(w http.ResponseWriter, code int, format string, args ...interface{}) {
	SendError(w, code, &gqlerror.Error{Message: fmt.Sprintf(format, args...)})
}

func toGQLError(err error) *gqlerror.Error {
	return &gqlerror.Error{
		Message: err.Error(),
	}
}
