package data

import (
	"errors"
	"strings"
)

type CaptorValue struct {
	Captor string
	Field  string
}

func (captorValue *CaptorValue) String() string {
	return captorValue.Captor + "." + captorValue.Field
}

func (captorValue *CaptorValue) FromString(s string) error {
	split := strings.Split(s, ".")
	if len(split) != 2 {
		return errors.New("invalid captor value")
	}
	captorValue.Captor = split[0]
	captorValue.Field = split[1]
	return nil
}
