package golang

import (
	// golang package
	"encoding/json"
)

// JsonMarshal returns the JSON encoding of input.
func (g *Golang) JsonMarshal(input interface{}) ([]byte, error) {
	return json.Marshal(input)
}

// JsonUnmarshal parses the JSON-encoded data and stores the result in the value pointed to by dest.
func (g *Golang) JsonUnmarshal(input []byte, dest interface{}) error {
	return json.Unmarshal(input, dest)
}
