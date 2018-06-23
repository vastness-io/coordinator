package errors

import "errors"

var (
	InvalidProject         = errors.New("invalid project")
	ProjectDoesNotExistErr = errors.New("can't find project")
	NotByteSliceErr = errors.New("unable to type assert to []byte")
)
