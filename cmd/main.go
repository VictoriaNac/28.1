package main

import (
    "github.com/VictoriaNac/28.1/interial/app"
    "github.com/VictoriaNac/28.1/interial/storage"
)

func main(){

//экземпляр структуры репозитория
	storage := storage.NewMemStorage() 
 //экземпляр структуры приложения
	app := &app.App{
		Storage:storage,
		}                            
	app.Run()
}