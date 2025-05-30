package bg_ui

import (
	"bgGoVisionneurVocabulaire/bg_metier"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

var numero int = 0
var word bg_metier.BgWord
var isFrenchDisplay = false

func MainUI(listWords []bg_metier.BgWord) error {

	fmt.Println("listWords lenxxx  :", len(listWords))

	word = listWords[0]
	myApp := app.New()
	myWindow := myApp.NewWindow("Interface Simple")

	checkbox := widget.NewCheck("French", func(checked bool) {
		if checked {
			isFrenchDisplay = true
		} else {
			isFrenchDisplay = false
		}
	})
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
		displayWord(labelNumero, labelInstruction, labelFrench, word)

	})

	buttonPrevious := widget.NewButton("Previous", func() {

		numero--
		if numero <= 0 {
			numero = len(listWords) - 1
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		displayWord(labelNumero, labelInstruction, labelFrench, word)
	})

	buttonTraduction := widget.NewButton("French", func() {

		labelFrench.SetText(word.LabelFr)
	})
	buttonAudioUK := widget.NewButton("Audio UK", func() {
		go playMP3(word.FilePathAudioUK)
	})
	buttonAudioAU := widget.NewButton("Audio AU", func() {
		go playMP3(word.FilePathAudioAU)
	})
	buttonAudioNeutre := widget.NewButton("Audio Neutre", func() {
		go playMP3(word.FilePathAudioUK2)
	})

	ligneMenu := container.NewHBox(
		checkbox, buttonAudioNeutre, buttonAudioUK, buttonAudioAU,
	)

	ligneHaut := container.NewHBox(
		labelNumero,
		labelInstruction,
	)

	ligneNextPrevious := container.NewGridWithColumns(2,
		buttonPrevious,
		buttonNext,
	)

	contenu := container.NewVBox(
		ligneMenu,
		ligneHaut,
		ligneNextPrevious,
		buttonTraduction,
		labelFrench,
	)

	myWindow.SetContent(contenu)
	myWindow.ShowAndRun()

	return nil
}

func displayWord(labelNumero *widget.Label, labelInstruction *widget.Label, labelFrench *widget.Label, word bg_metier.BgWord) {
	labelNumero.SetText(" " + strconv.Itoa(numero))

	labelInstruction.SetText(word.LabelEn)
	go playMP3(word.FilePathAudio)

	if isFrenchDisplay {
		labelFrench.SetText(word.LabelFr)
	} else {
		labelFrench.SetText("")
	}

}

func playMP3(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		log.Fatal(err)
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.SetVolume(1.0)
	p.Play()

	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}
}
