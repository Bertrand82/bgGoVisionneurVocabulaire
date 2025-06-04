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
var listWords []bg_metier.BgWord
var word bg_metier.BgWord
var isFrenchDisplay = false
var isAutomatic = false
var isOrder1 = false

var labelNumero = widget.NewLabel("numero : " + strconv.Itoa(numero))
var labelEnglish = widget.NewLabel(" ")
var labelFrench = widget.NewLabel(" ")

var checkboxDisplayIsAutomatic = widget.NewCheck("Auto", func(checked bool) {
	if checked {
		isAutomatic = true
	} else {
		isAutomatic = false
	}
	go playMP3All()
})

var checkboxDisplayFrench = widget.NewCheck("French", func(checked bool) {
	if checked {
		isFrenchDisplay = true
	} else {
		isFrenchDisplay = false
	}
})

var checkboxOrder1 = widget.NewCheck("Ordre 1", func(checked bool) {
	if checked {
		isOrder1 = true
	} else {
		isOrder1 = false
	}
})

func MainUI(listWordsArg []bg_metier.BgWord) error {
	listWords = listWordsArg
	fmt.Println("listWords lenxxx  :", len(listWords))

	word = listWords[0]
	myApp := app.New()
	myWindow := myApp.NewWindow("Interface Simple")

	buttonNext := widget.NewButton("Next", func() {

		numero--
		if numero < 0 {
			numero = len(listWords) - 1
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		displayWord(word)

	})

	buttonPrevious := widget.NewButton("Previous", func() {

		numero++
		if numero >= len(listWords) {
			numero = 0
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		displayWord(word)
	})
	buttonRepeat := widget.NewButton("Repeat", func() {

		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		displayWord(word)
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
		go playMP3(word.FilePathAudio_US_1)
	})
	buttonAudioUS2 := widget.NewButton("Audio US ", func() {
		go playMP3(word.FilePathAudio_US_2)
	})

	ligneMenu := container.NewHBox(
		checkboxOrder1, checkboxDisplayFrench, checkboxDisplayIsAutomatic, buttonAudioNeutre, buttonAudioUK, buttonAudioAU, buttonAudioUS2,
	)

	ligneHaut := container.NewHBox(
		labelNumero,
		labelEnglish,
	)

	ligneNextPrevious := container.NewGridWithColumns(3,
		buttonPrevious,
		buttonRepeat,
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

func displayWord(word bg_metier.BgWord) {
	labelNumero.SetText(" " + strconv.Itoa(numero))
	labelEnglish.SetText(word.LabelEn)
	isAutomatic = false
	checkboxDisplayIsAutomatic.SetChecked(false)
	go playMP3(word.FilePathAudio)

	if isFrenchDisplay {
		labelFrench.SetText(word.LabelFr)
	} else {
		labelFrench.SetText("")
	}

}
func playMP3All() {
	for isAutomatic {
		numero--
		if numero < 0 {
			numero = len(listWords) - 1
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		labelEnglish.SetText(word.LabelEn)
		labelFrench.SetText(word.LabelFr)
		playMP3(word.FilePathAudioUK)
		time.Sleep(100 * time.Millisecond)
	}
}

func playMP3(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
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
