package fileutils

import (
	"bgGoVisionneurVocabulaire/bg_metier"
	"fmt"
	"log"
)

var ListWords []bg_metier.BgWord

func Lire_fichier_by_name(nomFichier string) []bg_metier.BgWord {
	// Je reinitialise la liste
	ListWords = []bg_metier.BgWord{}
	err := LireFichierLignes(nomFichier, readLigne1)
	fmt.Println("listWords len:", len(ListWords))
	if err != nil {
		log.Fatalf("Une erreur est survenue : %v", err)
		return nil
	}
	return ListWords
}

func readLigne1(ligne string) {
	var word bg_metier.BgWord = bg_metier.NewBgBgWord(ligne)
	ListWords = append(ListWords, word)
	fmt.Println("Ligne luezzz word :", word.LabelEn)
}
