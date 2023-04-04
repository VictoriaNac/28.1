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



	/*
a := student.NewStudent("Вася", 24, 1)
	b := student.NewStudent("Семен", 32, 2)
	ss := storage.NewStudentStorage()
	ss.Put(a)
	ss.Put(b)
	fmt.Println(ss.Get(a.Name))
	fmt.Println(ss.Get(b.Name))
}
	*/
}