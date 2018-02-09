package myfunc

func NewInt() *int{
	return new(int)
}


func NewIntVar() *int{
	var dummy int
	return &dummy
}


func delta(old, new int )int{
	return new-old
}