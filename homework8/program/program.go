package program

import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/alextonkonogov/gb-golang-level-2/homework8/config"
	f "github.com/alextonkonogov/gb-golang-level-2/homework8/files"
)

type Program struct{}

func (p Program) Start(cnfg *config.AppConfig, uniqueFiles *f.UniqueFiles) error {
	fmt.Printf("Program starts searching for duplicate files in \"%s\"...\n", cnfg.Path)

	dublicates := 0
	files := make(chan f.File)

	go func(dir string, files chan<- f.File) {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
			}
			if !info.IsDir() && info.Name() != ".DS_Store" {
				files <- f.NewFile(path, info.Name())
			}
			return nil
		})
		close(files)
	}(cnfg.Path, files)

	var wg sync.WaitGroup
	wg.Add(cnfg.Workers)

	for i := 0; i < cnfg.Workers; i++ {
		func(files <-chan f.File, uniqueFiles *f.UniqueFiles, wg *sync.WaitGroup) {
			for file := range files {
				data, err := ioutil.ReadFile(path.Join(".", file.Path))
				if err != nil {
					log.Fatal(err)
				}
				digest := sha512.Sum512(data)
				uniqueFiles.Mtx.Lock()
				if _, ok := uniqueFiles.Map[digest]; ok {
					dublicates++
				}
				uniqueFiles.Map[digest] = append(uniqueFiles.Map[digest], file)
				uniqueFiles.Mtx.Unlock()
			}
			wg.Done()
		}(files, uniqueFiles, &wg)
	}

	wg.Wait()
	uniqueFiles.Sort()
	fmt.Printf("Found %d unique files and %d dublicates:\n", len(uniqueFiles.Map), dublicates)

	for k, _ := range uniqueFiles.Map {
		for i, _ := range uniqueFiles.Map[k] {
			if i == 0 {
				fmt.Println(uniqueFiles.Map[k][i].Name)
				if len(uniqueFiles.Map[k]) > 1 {
					fmt.Printf("    %d dublicates:\n", len(uniqueFiles.Map[k])-1)
				}
			} else {
				fmt.Printf("    %s\n", uniqueFiles.Map[k][i].Name)
			}
		}
	}

	if !cnfg.DeleteDublicates {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Are you sure to delete all duplicate files? yes/no: ")
	for scanner.Scan() {
		if scanner.Err() != nil {
			return scanner.Err()
		}
		in := strings.TrimSpace(scanner.Text())
		if in != "yes" && in != "no" {
			fmt.Printf("%s", "    try again: type yes or no ")
			continue
		}
		if in != "yes" {
			break
		}

		err := uniqueFiles.DeleteDuplicates()
		if err != nil {
			return err
		}
		fmt.Print("All duplicate files were deleted\n")
		break
	}

	return nil
}
