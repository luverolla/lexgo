package errs

type NotFoundErr struct{}

func NotFound() NotFoundErr {
	return NotFoundErr{}
}

func (err NotFoundErr) Error() string {
	return "Not found"
}
