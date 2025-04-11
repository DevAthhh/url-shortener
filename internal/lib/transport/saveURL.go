package transport

import "errors"

type RequestToSave struct {
	Root string `json:"root"`
	Size int    `json:"size"`
}

type ResponseFromSave struct {
	Status string `json:"status"`
	Alias  string `json:"alias"`
}

func (r *RequestToSave) Validate() error {
	if len(r.Root) == 0 || r.Size < 5 || r.Size > 100 {
		return errors.New("some parametr invalid")
	}
	return nil
}
