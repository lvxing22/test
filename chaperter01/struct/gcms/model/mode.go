package model

type Student struct{
	Name  string
}
type student1 struct{
	Name1  string
}
func Newstudent(n string) *student1{
	return &student1{
		Name1: n,
	}
}