package inspector

import "errors"

var (
	ErrNoConfig            = errors.New("no config provided")
	ErrUnknownTarget       = errors.New("unknown target provided")
	ErrUnknownInspector    = errors.New("unknown inspector")
	ErrNoGOPATH            = errors.New("no GOPATH variable found")
	ErrDstNotExists        = errors.New("destination directory doesn't exists")
	ErrUnknownEncodingType = errors.New("unknown encoding type")
	ErrUnsupportedType     = errors.New("unsupported type")
	ErrMustPointerType     = errors.New("variable must have pointer type")
)
