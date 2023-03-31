package main

import (
	"fmt"

	"28.1/interial/app"
	"28.1/interial/storage"
	sl "github.com/VictoriaNac/logg"
)

func main(){

	l, err :=sl.NewLogger ("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(l)



	storage := storage.NewStorage() //экземпляр структуры репозитория
	app := app.NewApp(storage)      //экземпляр структуры приложения
	app.Run()
}