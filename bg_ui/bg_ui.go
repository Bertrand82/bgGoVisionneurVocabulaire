package bg_ui

import (
	"bgGoVisionneurVocabulaire/bg_metier"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

var numero int = 0
var listWords2 []bg_metier.BgWord
var word bg_metier.BgWord
var isFrenchDisplay = false
var isAutomatic = false
var isOrder1 = false

type customTheme struct {
	fyne.Theme
}

func MainUI(listWords []bg_metier.BgWord) error {

	fmt.Println("listWords lenxxx  :", len(listWords))
	listWords2 = listWords
	word = listWords[0]
	myApp := app.New()

	//myApp.Settings().SetTheme(customTheme{})

	myWindow := myApp.NewWindow("Interface Simple")
	labelNumero := widget.NewLabel("numero : " + strconv.Itoa(numero))
	labelEnglish := widget.NewLabel(word.LabelEn + "  xxxxxxxxxxxxxxxxxxxxyyyyyyyyyyyyyyyyyyyy")
	labelFrench := widget.NewLabel("xxx")

	checkboxDisplayIsAutomatic := widget.NewCheck("Auto", func(checked bool) {
		if checked {
			isAutomatic = true
		} else {
			isAutomatic = false
		}
		go playMP3All(labelNumero, labelEnglish, labelFrench)
	})
	checkboxDisplayFrench := widget.NewCheck("French", func(checked bool) {
		if checked {
			isFrenchDisplay = true
		} else {
			isFrenchDisplay = false
		}

	})

	checkboxOrder1 := widget.NewCheck("Ordre 1", func(checked bool) {
		if checked {
			isOrder1 = true
		} else {
			isOrder1 = false
		}
	})

	buttonNext := widget.NewButton("Next", func() {

		numero--
		if numero < 0 {
			numero = len(listWords) - 1
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		displayWord(labelNumero, labelEnglish, labelFrench, word)

	})

	buttonPrevious := widget.NewButton("Previous", func() {

		numero++
		if numero >= len(listWords) {
			numero = 0
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords[numero]
		word = nextWord
		displayWord(labelNumero, labelEnglish, labelFrench, word)
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
func playMP3All(labelNumero *widget.Label, labelEnglish *widget.Label, labelFrench *widget.Label) {
	for isAutomatic {
		numero--
		if numero < 0 {
			numero = len(listWords2) - 1
		}
		labelNumero.SetText(" " + strconv.Itoa(numero))
		var nextWord = listWords2[numero]
		word = nextWord
		labelEnglish.SetText(word.LabelEn)
		labelFrench.SetText(word.LabelFr)
		playMP3(word.FilePathAudioUK)
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

func (m customTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return theme.DefaultTheme().Size(name) * 1.2 // Augmente la taille de 20%
	}
	return theme.DefaultTheme().Size(name)
}
