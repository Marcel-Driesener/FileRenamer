package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Willkommen zum File Renamer\n------------------")
	fmt.Println("Hier kannst du deine Dateien alle auf einmal umbenennen.\nWähle zwischen 2 Optionen: Alle Dateien im Ordner oder Welche mit bestimmten Namen")
	fmt.Println("\nImportant!")
	fmt.Println("Die Dateien werden OHNE die Nummern und dem Dateiformat angegeben (z.B 'Birthday-1.txt' wird zu 'Birthday-')\n es werden automatisch alle Dateien ausgewählt, die gleich heißen.\n Das selbe gilt für den neuen Namen, es wird automatisch alle Dateien gleich benannt, und nur mit einer Nummer unterschieden (Birthday_1,Birthday_2,...)")

	path := bufio.NewReader(os.Stdin)
	fmt.Print("Gib den Pfad (z.B. C:\\User\\Fotos) zum Ordner an: ")
	dir, _ := path.ReadString('\n')
	dir = strings.TrimSpace(dir)

	nfn := bufio.NewReader(os.Stdin)
	fmt.Print("Gib den neuen namen an wie die Dateien heißen solllen: ")
	newFileName, _ := nfn.ReadString('\n')
	newFileName = strings.TrimSpace(newFileName)

	fmt.Println("Sollen nur Dateien mit einem bestimmten Namen umbenannt werden?: ")
	ent := bufio.NewReader(os.Stdin)
	fmt.Println("Yes / No (alle dateien): ")
	entscheidung, _ := ent.ReadString('\n')
	entscheidung = strings.TrimSpace(entscheidung)
	entscheidung = strings.ToUpper(entscheidung)

	switch entscheidung {
	case "NO", "N":
		findAllFiles(dir, newFileName)
	case "YES", "Y":
		findSpecificFiles(dir, newFileName)
	default:
		fmt.Println("Invalid Characters")
	}

}

func findAllFiles(dir, newFileName string) {

	fmt.Println("Alle Dateien (ohne unterordner) werden umbenannt.")

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 1
	for _, file := range files {
		if !file.IsDir() {
			oldPath := filepath.Join(dir, file.Name())
			ext := filepath.Ext(file.Name())
			newPath := filepath.Join(dir, newFileName+strconv.Itoa(i)+ext)
			err = os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println(err)
			}
			i++
		}
	}
	fmt.Println("Dateien wurden umbenannt.")
}

func findSpecificFiles(dir, newFileName string) {

	ofn := bufio.NewReader(os.Stdin)
	fmt.Print("Gib den Namen (z.B Birthday-) der Dateien an die umbenannt werden soll: ")
	oldFileName, _ := ofn.ReadString('\n')
	oldFileName = strings.TrimSpace(oldFileName)

	pattern := oldFileName + "\\d{19}.*"
	re := regexp.MustCompile(pattern)

	// Walk through the directory tree
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a regular file and if it matches the pattern
		if !info.IsDir() && re.MatchString(info.Name()) {
			fmt.Println(info.Name())
		}
		i := 1
		if strings.Contains(info.Name(), oldFileName) {
			ext := filepath.Ext(info.Name())
			err := os.Rename(path, filepath.Join(filepath.Dir(path), newFileName+strconv.Itoa(i)+ext))
			if err != nil {
				fmt.Println("Fehler beim umbenennen:", err)
				return err
			}
			i++
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
