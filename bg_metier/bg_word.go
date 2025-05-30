package bg_metier

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"

	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// Définition de la structure
type BgWord struct {
	LabelEn          string
	LabelFr          string
	Numero           int
	Coef             int
	FileNameAudio    string
	FilePathAudio    string
	FilePathAudioUK  string
	FilePathAudioAU  string
	FilePathAudioUK2 string
}

const tempDir string = "AudioTemp"
const tempDirAudio string = tempDir + "\\" + "Neutre"
const tempDirAudioUK string = tempDir + "\\" + "UK"
const tempDirAudioUK2 string = tempDir + "\\" + "UK_2"
const tempDirAudioAU string = tempDir + "\\" + "AU"

func NewBgBgWord(line string) BgWord {
	before, after, findToken := strings.Cut(line, ":")
	// Initialisation d'une instance de la structure
	if findToken {

		var fileName string = "audio_" + strings.ReplaceAll(strings.TrimSpace(before), " ", "_") + "_x_"

		var pathEnglishUK_2 = recordMP3_hegedustibor(tempDirAudioUK, fileName, before, voices.EnglishUK)
		var pathEnglishAU_2 = recordMP3_hegedustibor(tempDirAudioAU, fileName, before, voices.EnglishAU)
		var pathEnglish_2 = recordMP3_hegedustibor(tempDirAudio, fileName, before, voices.English)
		var pathEnglish_3 = recordMP3_googleAPI(tempDirAudioUK2, fileName, before, "en_GB", "en-GB-Wavenet-D")

		objet := BgWord{
			LabelEn:          before,
			LabelFr:          after,
			Numero:           1,
			Coef:             5,
			FileNameAudio:    fileName,
			FilePathAudio:    pathEnglish_2,
			FilePathAudioUK:  pathEnglishUK_2,
			FilePathAudioAU:  pathEnglishAU_2,
			FilePathAudioUK2: pathEnglish_3,
		}

		// Affichage des champs de l'objet
		fmt.Println("Numero:", objet.Numero)
		fmt.Println("LabelEn:", "   "+objet.LabelEn+"     ::::::: "+objet.LabelFr)

		fmt.Println("Coef:", objet.Coef)
		fmt.Println("Coef:", objet.FilePathAudio)
		return objet
	} else {
		objet := BgWord{
			LabelEn: line,
			LabelFr: "",
			Numero:  1,
			Coef:    5,
		}
		return objet
	}
}

func recordMP3_hegedustibor(dirName string, fileName string, before string, langage string) string {
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
func recordMP3_googleAPI(dirName string, fileName string, word string, langage string, voix string) string {
	var pathEnglishUK = filepath.Join(dirName, fileName+".mp3")
	var fileExistEnglishUK bool

	if _, err := os.Stat(pathEnglishUK); err == nil {
		fileExistEnglishUK = true
	} else if os.IsNotExist(err) {
		fileExistEnglishUK = false
	}
	if !fileExistEnglishUK {
		err := SynthesizeToFileByGoogleAPI(word, fileName+".mp3", dirName, langage, voix)
		if err != nil {
			fmt.Printf("Erreur : %v\n", err)
		}
	}

	fmt.Println("langage ", langage+"  voix"+voix+"   path:"+pathEnglishUK)
	return pathEnglishUK
}

func SynthesizeToFileByGoogleAPI(text, filename, dir, languageCode, voiceName string) error {
	ctx := context.Background()

	// Crée un client Text-to-Speech
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("échec de la création du client Text-to-Speech : %v", err)
	}
	defer client.Close()

	// Prépare la requête de synthèse vocale
	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: languageCode,
			Name:         voiceName,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	// Effectue la requête de synthèse vocale
	resp, err := client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return fmt.Errorf("échec de la synthèse vocale : %v", err)

	}

	// Crée le répertoire de destination s'il n'existe pas
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("échec de la création du répertoire %s : %v", dir, err)
	}

	// Construit le chemin complet du fichier de sortie
	outputPath := filepath.Join(dir, filename)

	// Écrit le contenu audio dans le fichier
	if err := ioutil.WriteFile(outputPath, resp.AudioContent, 0644); err != nil {
		return fmt.Errorf("échec de l'écriture du fichier audio : %v", err)
	}

	fmt.Printf("Le fichier audio a été écrit avec succès : %s\n", outputPath)
	return nil
}
