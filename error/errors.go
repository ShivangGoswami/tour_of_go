package main

type MyError struct{
	When time.time
	What string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s",e.When,e.What)
}

