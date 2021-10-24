package program_test

import (
	"log"
	"testing"

	"github.com/alextonkonogov/gb-golang-level-2/homework8/config"
	"github.com/alextonkonogov/gb-golang-level-2/homework8/files"
	"github.com/alextonkonogov/gb-golang-level-2/homework8/program"
)

func Example() {
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	uniqueFiles := files.NewUniqueFilesMap()

	program := program.NewProgram(cnfg, uniqueFiles)
	err = program.Start()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	//Program starts searching for duplicate files in "."...
	//Found 2 unique files and 0 duplicates:
	//program.go
	//program_test.go
}

func BenchmarkProgram_Start(b *testing.B) {
	for j := 0; j < b.N; j++ {
		cnfg, err := config.NewAppConfig()
		if err != nil {
			log.Fatal(err)
		}
		cnfg.PrintResult = false
		uniqueFiles := files.NewUniqueFilesMap()

		program := program.NewProgram(cnfg, uniqueFiles)
		err = program.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}

//goos: darwin
//goarch: amd64
//pkg: github.com/alextonkonogov/gb-golang-level-2/homework8/program
//BenchmarkProgram_Start
//BenchmarkProgram_Start-8   	   23029	     49370 ns/op
//PASS
