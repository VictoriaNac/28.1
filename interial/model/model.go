//1.описание данных с которыми мы будем работать


package model  //старое название файла  "student"

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