package bg_ui

import (
	"bgGoVisionneurVocabulaire/bg_metier"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var numero int = 0
var word bg_metier.BgWord

func MainUI(listWords []bg_metier.BgWord) error {

	fmt.Println("listWords lenxxx  :", len(listWords))
	word = listWords[0]
	myApp := app.New()
	myWindow := myApp.NewWindow("Interface Simple")

	labelNumero := widget.NewLabel("numero : " + strconv.Itoa(numero))
	labelInstruction := widget.NewLabel(word.LabelEn + "  xxxxxxxxxxxxxxxxxxxxyyyyyyyyyyyyyyyyyyyy")
	labelFrench := widget.NewLabel("xxx")

	buttonNext := widget.NewButton("Next", func() {

		numero++
		if numero >= len(listWords) {
			numero = 0
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		labelInstruction.SetText(word.LabelEn)
		labelFrench.SetText("")
	})

	buttonPrevious := widget.NewButton("Previous", func() {

		numero--
		if numero <= 0 {
			numero = len(listWords) - 1
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		labelInstruction.SetText(word.LabelEn)
		labelFrench.SetText("")
	})

	buttonTraduction := widget.NewButton("French", func() {

		labelFrench.SetText(word.LabelFr)
	})

	ligneHaut := container.NewHBox(
		labelNumero,
		labelInstruction,
	)

	ligneNextPrevious := container.NewGridWithColumns(2,
		buttonPrevious,
		buttonNext,
	)

	contenu := container.NewVBox(
		ligneHaut,
		ligneNextPrevious,
		buttonTraduction,
		labelFrench,
	)

	myWindow.SetContent(contenu)
	myWindow.ShowAndRun()

	return nil
}
