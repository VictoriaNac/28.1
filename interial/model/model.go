//1.описание данных с которыми мы будем работать


package model  

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age, grade int) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}