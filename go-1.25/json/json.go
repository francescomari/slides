package json

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"strings"
)

type Gopher struct {
	Name string `json:"name"`
}

func ReaderWriter() error {
	// BEGIN JSON RW OMIT
	var buffer strings.Builder
	if err := json.MarshalWrite(&buffer, Gopher{Name: "Bob"}); err != nil { // HL
		return err
	}
	var gopher Gopher
	if err := json.UnmarshalRead(strings.NewReader(buffer.String()), &gopher); err != nil { // HL
		return err
	}
	// END JSON RW OMIT
	return nil
}

func EncoderDecoder() error {
	// BEGIN JSON ENC OMIT
	var buffer strings.Builder
	encoder := jsontext.NewEncoder(&buffer) // HL
	if err := json.MarshalEncode(encoder, Gopher{Name: "Bob"}); err != nil {
		return err
	}
	var gopher Gopher
	decoder := jsontext.NewDecoder(strings.NewReader(buffer.String())) // HL
	if err := json.UnmarshalDecode(decoder, &gopher); err != nil {
		return err
	}
	// END JSON ENC OMIT
	return nil
}
