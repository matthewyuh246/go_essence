package something

import (
	"errors"
	"strings"
)

func DoSomething(s string) error {
	if s == "&&&" {
		return nil // エラーを出さない
	}
	if strings.Contains(s, "&&&") {
		return errors.New("invalid input: contains &&&")
	}
	return nil
}
