package greetings

import (
	"fmt"
	"errors"
)


func Hello(name string)  (string, error) {
	//if no name was giving return an erroe with a message
	if name == "" {
		return "", errors.New("empty name !")
	}

	//if a name was giving return a message with the name

	message := fmt.Sprintf("Hello, %v. welcome", name)
	return message, nil

}
