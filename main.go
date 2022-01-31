package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()                   // Инициализация приложения
	w := a.NewWindow("Hello Fyne")   // Создание окна
	w.Resize(fyne.NewSize(400, 400)) // Задаем размер окна

	w.SetContent(widget.NewLabel("Hello World!")) // Наполняем контентом
	w.ShowAndRun()                                // Запуск окна
}
