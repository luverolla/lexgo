package errs

type NotFoundErr struct{}

func NotFound() NotFoundErr {
	return NotFoundErr{}
}

func (err NotFoundErr) Error() string {
	return "Not found"
}

type EmptyErr struct{}

func Empty() EmptyErr {
	return EmptyErr{}
}

func (err EmptyErr) Error() string {
	return "Attempted to Get/Peek/Pop/Remove from an empty collection"
}
