package main

import (
    "github.com/VictoriaNac/28.1/interial/app"
    "github.com/VictoriaNac/28.1/interial/storage"
)

func main(){

//экземпляр структуры репозитория
	storage := storage.NewStorage() 
 //экземпляр структуры приложения
	app := &app.App{
		Storage:storage,
		}                            
	app.Run()
	

//Реализуйте методы в репозитории и вывод списка зачисленных студентов

type Storage map[string]*student.Student

func (NewStorage) (Storage); {
	return make(Storage)
}

func (s Storage) Put(value *student.Student) error {
	s[value.GetName()] = value

	if s[value.GetName()] == nil {
		return errors.New("smt")
	} else {
		return nil
	}
}

func (s Storage) Get(name string) (*student.Student, error) {
	if s[name] == nil {
		return nil, errors.New("smt1")
	} else {
		return s[name], nil
	}
}

func (s *Storage) PrintStudents() {
	for _, value := range *s {
		fmt.Print(value.GetName(), " ")
		fmt.Print(value.GetAge(), " ")
		fmt.Print(value.GetGrade(), "\n")
	}
}

}