package durl

import (
	"encoding/base64"
	"errors"
	"regexp"
)

func DataUrlParser(durl string) ([]byte, error) {
	re := regexp.MustCompile(`^data:(.*?);base64,(.*?)$`)
	r := re.FindSubmatch([]byte(durl))

	if len(r) != 3 {
		return nil, errors.New("format error")
	}
	data, err := base64.StdEncoding.DecodeString(string(r[2][:]))
	return data, err
}
