package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type Op string

type Msg map[string]string

type Error struct {
	err      error
	status   int
	op       Op
	messages Msg
}

func Errorf(op Op, extra ...interface{}) error {
	err := Error{
		op:       op,
		status:   http.StatusInternalServerError,
		messages: Msg{},
	}

	for _, ex := range extra {
		switch t := ex.(type) {
		case error:
			err.err = t
		case int:
			err.status = t
		}
	}
	return err

}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Message() string {
	var parts []string
	for field, msg := range e.messages {
		parts = append(parts, fmt.Sprintf("%s:%s", field, msg))
	}
	return fmt.Sprintf("field errors:%s", strings.Join(parts, ","))
}
