package util

import (
	"errors"
	"github.com/google/uuid"
	"strings"
)

func GenerateId() string {
	return uuid.NewString()
}

func GetOwnerAndNameFromId(id string) (string, string) {
	tokens := strings.Split(id, "/")
	if len(tokens) != 2 {
		panic(errors.New("GetOwnerAndNameFromId() error, wrong token count for ID: " + id))
	}

	return tokens[0], tokens[1]
}
