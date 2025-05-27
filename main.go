package main

import (
	"bgGoVisionneurVocabulaire/bg_metier"
	"bgGoVisionneurVocabulaire/bg_ui"
	"bgGoVisionneurVocabulaire/fileutils"
	"fmt"
	"log"
)

var listWords []bg_metier.BgWord

func main() {
	nomFichier := "vocabulaire.txt"

	err := fileutils.LireFichierLignes(nomFichier, readLigne1)
	fmt.Println("listWords len:", len(listWords))
	if err != nil {
		log.Fatalf("Une erreur est survenue : %v", err)
		return
	} else {
		bg_ui.MainUI(listWords)
	}

}

func readLigne1(ligne string) {
	var word bg_metier.BgWord = bg_metier.NewBgBgWord(ligne)
	listWords = append(listWords, word)
	fmt.Println("Ligne luezzz word :", word.LabelEn)
}
