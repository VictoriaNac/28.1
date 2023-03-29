package main


import(
	"28.1/interial/app"
	"28.1/interial/storage"
)

func main(){
	storage := storage.NewStorage() //экземпляр структуры репозитория
	app := app.NewApp(storage)      //экземпляр структуры приложения
	app.Run()
}