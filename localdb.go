package mimetypes

import (
	"encoding/json"
	"fmt"

	_ "embed"
)

// Using the following github discussion:
//
// https://gist.github.com/Qti3e/6341245314bf3513abb080677cd1c93b

//go:embed mimetypes.json
var mimetypes []byte

var mimeTypes, _ = func() (m map[string]mimeType, err error) {
	err = json.Unmarshal(mimetypes, &m)
	return
}()

type mimeType struct {
	Mime  string `json:"mime"`
	Signs []sign `json:"signs"`
}

type sign struct {
	Offset int
	Bytes  []byte
}

func (s *sign) UnmarshalJSON(data []byte) error {
	var tmp string
	var err = json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	var (
		n      int
		offset int
		bytes  []byte
	)
	n, err = fmt.Sscanf(tmp, "%d,%x", &offset, &bytes)
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("invalid sign format")
	}
	s.Offset = offset
	s.Bytes = bytes
	return nil
}
