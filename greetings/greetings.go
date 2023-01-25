package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)


func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages :=make(map[string]string)
	for _, name := range names {
		message, error :=Hello(name)
		if(error!=nil) {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v welcome!",
		"Great to see you %v!",
		"Hail, %v well met!",
	}

	return formats[rand.Intn(len(formats))]
}