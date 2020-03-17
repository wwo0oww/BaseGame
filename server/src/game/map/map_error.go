package emap

import "errors"

func MapError(err string) error{
 	return errors.New(err)
}