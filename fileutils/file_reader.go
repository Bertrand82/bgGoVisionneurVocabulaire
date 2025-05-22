package fileutils

import (
	"bufio"
	"fmt"
	"os"
)

// LireFichierLignes lit un fichier texte ligne par ligne et applique une fonction de traitement à chaque ligne.
func LireFichierLignes(nomFichier string, traiterLigne func(string)) error {
	fichier, err := os.Open(nomFichier)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier %s : %v", nomFichier, err)
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		traiterLigne(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier %s : %v", nomFichier, err)
	}

	return nil
}
