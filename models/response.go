package models

import "fmt"

type Response struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}

func (r *Response) GoString() string {
	return fmt.Sprintf(`
{
	Status: %d,
	Message: %s,
}`,
		r.Status,
		r.Message,
	)
}