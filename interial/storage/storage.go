//3.хранение данных в памяти

package storage

import(
	"github.com/VictoriaNac/28.1/interial/model"
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