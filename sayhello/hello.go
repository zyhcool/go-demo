package sayhello

import (
	"errors"
	"fmt"
)

func SayIt(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty errors")
	}
	message := fmt.Sprintf("this is %v", s)
	return message, nil
}
