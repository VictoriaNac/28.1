//2.взаимодействие с пользователем

package app

import(
	"github.com/VictoriaNac/28.1/interial/storage"
	"github.com/VictoriaNac/28.1/interial/model"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type App struct{
	Storage *storage.MemStorage
}

func (app *App) Run() {
	var buffer = bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя возраст курс студента")

	for {
		fmt.Print(">>> ")
		line, err := buffer.ReadString('\n')

		if err == io.EOF {
			println("Студенты из хранилища")
			break
		}

		lineFieldsArr := strings.Fields(line)

		if len(lineFieldsArr) != 3 {
			fmt.Println("Некорректный ввод")
			continue
		}

		studentName := lineFieldsArr[0]
		studentAge, errAge := strconv.Atoi(lineFieldsArr[1])
		studentGrade, errGrade := strconv.Atoi(lineFieldsArr[2])

		if errAge != nil || errGrade != nil {
			fmt.Println("Ошибка обработки числовых значений")
			continue
		} 

		student := model.NewStudent(studentName, studentAge, studentGrade)
	  

		err = app.Storage.Put(student)
		if err != nil {
			
		} else {
			fmt.Println("Студент c таким именем существует")
		}
	}
}