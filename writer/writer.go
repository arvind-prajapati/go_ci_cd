package writer

import (
	"fmt"
	"net/http"
)

type Writer interface {
	error
	Message() string
}

func WriteResponse(w http.ResponseWriter, err error) {
	if cr, ok := err.(Writer); ok {
		message := cr.Message()
		fmt.Println("Message of error:%s", message)
	}
}