package main

import (
    "github.com/VictoriaNac/28.1/interial/app"
    "github.com/VictoriaNac/28.1/interial/storage"
)

func main(){

	storage := storage.NewStorage() //экземпляр структуры репозитория
	app := app.NewApp(storage)      //экземпляр структуры приложения
	app.Run()
}