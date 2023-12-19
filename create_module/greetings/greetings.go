package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Greets a given name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}


	message := fmt.Sprintf(randomGreetingFormat(), name)
	return message, nil
}

func randomGreetingFormat() string {
	var formats = []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail %v!",
	} 

	return formats[rand.Intn(len(formats))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	} 
	return messages, nil
}