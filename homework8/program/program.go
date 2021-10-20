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

type Program struct {
	Config      *config.AppConfig
	UniqueFiles *f.UniqueFiles
	Duplicates  int
}

func (p *Program) Start() error {
	fmt.Printf("Program starts searching for duplicate files in \"%s\"...\n", p.Config.Path)

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
	}(p.Config.Path, files)

	var wg sync.WaitGroup
	wg.Add(p.Config.Workers)

	for i := 0; i < p.Config.Workers; i++ {
		func(files <-chan f.File, uniqueFiles *f.UniqueFiles, wg *sync.WaitGroup) {
			for file := range files {
				data, err := ioutil.ReadFile(path.Join(".", file.Path))
				if err != nil {
					log.Fatal(err)
				}
				digest := sha512.Sum512(data)
				uniqueFiles.Mtx.Lock()
				if _, ok := uniqueFiles.Map[digest]; ok {
					p.Duplicates++
				}
				uniqueFiles.Map[digest] = append(uniqueFiles.Map[digest], file)
				uniqueFiles.Mtx.Unlock()
			}
			wg.Done()
		}(files, p.UniqueFiles, &wg)
	}
	wg.Wait()

	p.UniqueFiles.Sort()
	p.printResult()

	if p.Config.DeleteDublicates && p.Duplicates > 0 {
		err := p.askForConfirmBeforeDeletion()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Program) printResult() {
	fmt.Printf("Found %d unique files and %d dublicates:\n", len(p.UniqueFiles.Map), p.Duplicates)

	for k, _ := range p.UniqueFiles.Map {
		for i, _ := range p.UniqueFiles.Map[k] {
			if i == 0 {
				fmt.Println(p.UniqueFiles.Map[k][i].Name)
				if len(p.UniqueFiles.Map[k]) > 1 {
					fmt.Printf("    %d dublicates:\n", len(p.UniqueFiles.Map[k])-1)
				}
			} else {
				fmt.Printf("    %s\n", p.UniqueFiles.Map[k][i].Name)
			}
		}
	}
}

func (p *Program) askForConfirmBeforeDeletion() error {
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

		err := p.UniqueFiles.DeleteDuplicates()
		if err != nil {
			return err
		}
		fmt.Print("All duplicate files were deleted\n")
		break
	}
	return nil
}

func NewProgram(cnfg *config.AppConfig, uniqueFiles *f.UniqueFiles) *Program {
	return &Program{cnfg, uniqueFiles, 0}
}
