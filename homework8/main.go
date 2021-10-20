package main

import (
	"log"

	"github.com/alextonkonogov/gb-golang-level-2/homework8/config"
	f "github.com/alextonkonogov/gb-golang-level-2/homework8/files"
	p "github.com/alextonkonogov/gb-golang-level-2/homework8/program"
)

func main() {
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	uniqueFiles := f.NewUniqueFilesMap()

	program := p.NewProgram(cnfg, uniqueFiles)
	err = program.Start()
	if err != nil {
		log.Fatal(err)
	}
}
