//go:build js
// +build js

package clipboard

import (
	"errors"
)

func readAll() (string, error) {
	return "", errors.New("not implemented")
}

func writeAll(text string) error {
	return errors.New("not implemented")
}
