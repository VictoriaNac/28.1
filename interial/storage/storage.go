//3.хранение данных в памяти

package storage

import(
	"github.com/VictoriaNac/28.1/interial/model"
	"fmt"
)

type MemStorage struct{
	studentByName map[string] *model.Student
}



func NewStorage() *MemStorage {
	return nil
}

func (s *MemStorage) Put(student *model.Student) error {
	return nil
}
func (s *MemStorage) Get(studentName string)(*model.Student, error) {
	return nil, nil
}
func (s *MemStorage) GetAll() []model.Student {
	return nil
}


//Реализуйте методы в репозитории и вывод списка зачисленных студентов


func (s *student) PrintStudents() {
	for _, value := range *s {
		fmt.Print(value.GetName(), " ")
		fmt.Print(value.GetAge(), " ")
		fmt.Print(value.GetGrade(), "\n")
	}
}