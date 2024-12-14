package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Willkommen zum File Renamer\n------------------")
	fmt.Println("Hier kannst du deine Dateien alle auf einmal umbenennen.")
	fmt.Println("Wähle zwischen 2 Optionen: Alle Dateien im Ordner (inklusive Unterordner) oder nur bestimmte Dateien")
	fmt.Println("\nHinweis:")
	fmt.Println("Die Dateien werden ohne die Nummern und das Dateiformat angegeben (z.B. 'Birthday-1.txt' wird zu 'Birthday-').")
	fmt.Println("Automatisch werden alle Dateien ausgewählt, die gleich heißen.")
	fmt.Println("Beim neuen Namen werden alle Dateien gleich benannt und nur durch eine Nummer unterschieden (z.B. 'Birthday_1', 'Birthday_2', ...)")

	dir := getInput("Gib den Pfad zum Ordner an (z.B. C:\\User\\Fotos): ")
	newFileName := getInput("Gib den neuen Namen an, wie die Dateien heißen sollen: ")

	choice := getInput("Sollen nur Dateien mit einem bestimmten Namen umbenannt werden? (yes/no): ")
	choice = strings.ToLower(choice)

	switch choice {
	case "no", "n":
		renameAllFiles(dir, newFileName)
	case "yes", "y":
		renameSpecificFiles(dir, newFileName)
	default:
		fmt.Println("Ungültige Eingabe. Das Programm wird beendet.")
	}
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func renameAllFiles(dir, newFileName string) {
	fmt.Println("Alle Dateien (inklusive Unterordner) werden umbenannt...")

	i := 1
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Fehler beim Zugriff auf Datei '%s': %v\n", path, err)
			return err
		}
		if !d.IsDir() {
			ext := filepath.Ext(d.Name())
			newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf("%s%d%s", newFileName, i, ext))

			if err := os.Rename(path, newPath); err != nil {
				fmt.Printf("Fehler beim Umbenennen von '%s': %v\n", d.Name(), err)
			} else {
				i++
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Fehler beim Durchsuchen des Verzeichnisses: %v\n", err)
	} else {
		fmt.Println("Dateien wurden erfolgreich umbenannt.")
	}
}

func renameSpecificFiles(dir, newFileName string) {
	oldFileName := getInput("Gib den Namen der Datei an, die umbenannt werden soll: ")
	pattern := fmt.Sprintf("^%s.*", regexp.QuoteMeta(oldFileName))
	re := regexp.MustCompile(pattern)

	fmt.Println("Die entsprechenden Dateien werden umbenannt...")

	i := 1
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Fehler beim Zugriff auf Datei '%s': %v\n", path, err)
			return err
		}
		if !d.IsDir() && re.MatchString(d.Name()) {
			ext := filepath.Ext(d.Name())
			newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf("%s%d%s", newFileName, i, ext))

			if err := os.Rename(path, newPath); err != nil {
				fmt.Printf("Fehler beim Umbenennen von '%s': %v\n", d.Name(), err)
			} else {
				i++
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Fehler beim Durchsuchen des Verzeichnisses: %v\n", err)
	} else {
		fmt.Println("Dateien wurden erfolgreich umbenannt.")
	}
}
