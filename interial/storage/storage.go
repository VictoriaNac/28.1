//3.хранение данных в памяти

package storage

import(
	"github.com/VictoriaNac/28.1/interial/model"
)

type MemStorage struct{
	studentByName map[string] *model.Student
}

var _ app.Storage = (*MemStorage)(nil)

func NewMemStorage() *MemStorage {
	return &MemStorage{
		studentByName: make(map[string]*model.Student),
	}
}

func (ms *MemStorage) GetAll() ([] *model.Student, error){
	var student []*model.Student
	for _, v :=range ms.studentByName {
		student = append (student, v)
	}
	return student, nil
}

func (ms *MemStorage) Put(student *model.Student) error {
	return nil
}

/*func NewStorage() *MemStorage {
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
}*/