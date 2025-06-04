package bg_ui

import (
	"bgGoVisionneurVocabulaire/bg_metier"
	"bgGoVisionneurVocabulaire/fileutils"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
var isAudioUK = false
var checkboxAudioUK = widget.NewCheck("Audio UK", func(checked bool) {
	if checked {
		isAudioUK = true
	} else {
		isAudioUK = false
	}
})
var isAudioAU = false
var checkboxAudioAU = widget.NewCheck("Audio AU", func(checked bool) {
	if checked {
		isAudioAU = true
	} else {
		isAudioAU = false
	}
})

var isAudioNeutre = false
var checkboxAudioNeutre = widget.NewCheck("Neutre", func(checked bool) {
	if checked {
		isAudioNeutre = true
	} else {
		isAudioNeutre = false
	}
})
var isAudioUS2 = false
var checkboxAudioUS2 = widget.NewCheck("US", func(checked bool) {
	if checked {
		isAudioUS2 = true
	} else {
		isAudioUS2 = false
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
		go playMP3File(word.FilePathAudioUK)
	})
	buttonAudioAU := widget.NewButton("Audio AU", func() {
		go playMP3File(word.FilePathAudioAU)
	})
	buttonAudioNeutre := widget.NewButton("Audio Neutre", func() {
		go playMP3File(word.FilePathAudio_US_1)
	})
	buttonAudioUS2 := widget.NewButton("Audio US ", func() {
		go playMP3File(word.FilePathAudio_US_2)
	})

	buttonChooseFile := widget.NewButton("Ouvrir un fichier", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println("Erreur :", err)
				return
			}
			if reader == nil {
				fmt.Println("Aucun fichier sélectionné")
				return
			}
			fmt.Println("Fichier sélectionné :", reader.URI().Path())
			readFile(reader)
			reader.Close()
		}, myWindow)
	})

	ligneMenu1 := container.NewHBox(
		checkboxOrder1, checkboxDisplayFrench, checkboxDisplayIsAutomatic, buttonAudioNeutre, buttonAudioUK, buttonAudioAU, buttonAudioUS2, buttonChooseFile,
	)
	ligneMenu2 := container.NewHBox(
		checkboxAudioNeutre, checkboxAudioUK, checkboxAudioAU, checkboxAudioUS2)

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
		ligneMenu1,
		ligneMenu2,
		ligneHaut,
		ligneNextPrevious,
		buttonTraduction,
		labelFrench,
	)

	myWindow.SetContent(contenu)
	myWindow.ShowAndRun()

	return nil
}

func readFile(reader fyne.URIReadCloser) {
	fmt.Println("Fichier sélectionné222 :", reader.URI().Path())
	listWords = fileutils.Lire_fichier_by_name(reader.URI().Path())
}
func displayWord(word bg_metier.BgWord) {
	labelNumero.SetText(" " + strconv.Itoa(numero))
	labelEnglish.SetText(word.LabelEn)
	isAutomatic = false
	checkboxDisplayIsAutomatic.SetChecked(false)
	go playMP3(word)

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
		playMP3(word)
		time.Sleep(100 * time.Millisecond)
	}
}
func playMP3(word bg_metier.BgWord) {

	if isAudioNeutre {
		playMP3File(word.FilePathAudio_US_1)
	}
	if isAudioUK {
		playMP3File(word.FilePathAudioUK)
	}
	if isAudioAU {
		playMP3File(word.FilePathAudioAU)
	}
	if isAudioUS2 {
		playMP3File(word.FilePathAudio_US_2)
	}
}
func playMP3File(filename string) {

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
