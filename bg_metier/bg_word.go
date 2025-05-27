package bg_metier

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

// Définition de la structure
type BgWord struct {
	LabelEn         string
	LabelFr         string
	Numero          int
	Coef            int
	FileNameAudio   string
	FilePathAudio   string
	FilePathAudioUK string
	FilePathAudioAU string
}

const tempDir string = "AudioTemp"
const tempDirAudio string = tempDir + "\\" + "Neutre"
const tempDirAudioUK string = tempDir + "\\" + "UK"
const tempDirAudioAU string = tempDir + "\\" + "AU"

func NewBgBgWord(line string) BgWord {
	before, after, findToken := strings.Cut(line, ":")
	// Initialisation d'une instance de la structure
	if findToken {

		var fileName string = "audio_" + strings.ReplaceAll(strings.TrimSpace(before), " ", "_") + "_x_"

		var pathEnglishUK_2 = recordMP3(tempDirAudioUK, fileName, before, voices.EnglishUK)
		var pathEnglishAU_2 = recordMP3(tempDirAudioAU, fileName, before, voices.EnglishAU)
		var pathEnglish_2 = recordMP3(tempDirAudio, fileName, before, voices.English)

		objet := BgWord{
			LabelEn:         before,
			LabelFr:         after,
			Numero:          1,
			Coef:            5,
			FileNameAudio:   fileName,
			FilePathAudio:   pathEnglish_2,
			FilePathAudioUK: pathEnglishUK_2,
			FilePathAudioAU: pathEnglishAU_2,
		}

		// Affichage des champs de l'objet
		fmt.Println("LabelEn:", objet.LabelEn)
		fmt.Println("LabelFr:", objet.LabelFr)
		fmt.Println("Numero:", objet.Numero)
		fmt.Println("Coef:", objet.Coef)
		fmt.Println("Coef:", objet.FilePathAudio)
		return objet
	} else {
		objet := BgWord{
			LabelEn: before,
			LabelFr: "",
			Numero:  1,
			Coef:    5,
		}
		return objet
	}
}

func recordMP3(dirName string, fileName string, before string, langage string) string {
	var pathEnglishUK = filepath.Join(dirName, fileName+".mp3")
	var fileExistEnglishUK bool
	if _, err := os.Stat(pathEnglishUK); err == nil {
		fileExistEnglishUK = true
	} else if os.IsNotExist(err) {
		fileExistEnglishUK = false
	}
	if !fileExistEnglishUK {
		speechEnglishUK := htgotts.Speech{
			Folder:   dirName, // Dossier de destination pour le fichier audio
			Language: langage, // Langue de synthèse vocale
		}
		path, err := speechEnglishUK.CreateSpeechFile(before, fileName)
		if err != nil {
			fmt.Println("ERRROR TTS bg", err)
		}
		pathEnglishUK = path
	}
	return pathEnglishUK
}
