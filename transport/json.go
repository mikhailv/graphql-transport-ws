package transport

import (
	"bytes"
	"encoding/json"
	"io"
)

func jsonDecodeReader(r io.Reader, val interface{}) error {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	return dec.Decode(val)
}

func jsonDecode(b []byte, val interface{}) error {
	return jsonDecodeReader(bytes.NewReader(b), val)
}

func jsonEncode(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}
