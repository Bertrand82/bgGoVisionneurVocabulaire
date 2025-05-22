package bg_metier

import (
	"fmt"
	"strings"
)

// Définition de la structure
type BgWord struct {
	LabelEn string
	LabelFr string
	Numero  int
	Coef    int
}

func NewBgBgWord(line string) BgWord {
	before, after, findToken := strings.Cut(line, ":")
	// Initialisation d'une instance de la structure
	if findToken {
		objet := BgWord{
			LabelEn: before,
			LabelFr: after,
			Numero:  1,
			Coef:    5,
		}

		// Affichage des champs de l'objet
		fmt.Println("LabelEn:", objet.LabelEn)
		fmt.Println("LabelFr:", objet.LabelFr)
		fmt.Println("Numero:", objet.Numero)
		fmt.Println("Coef:", objet.Coef)
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
