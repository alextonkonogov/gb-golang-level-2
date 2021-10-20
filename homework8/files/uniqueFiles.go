package files

import (
	"crypto/sha512"
	"fmt"
	"os"
	"sort"
	"sync"
)

type UniqueFiles struct {
	Mtx *sync.Mutex
	Map map[[sha512.Size]byte][]File
}

func (uf *UniqueFiles) Sort() {
	for k, _ := range uf.Map {
		if len(uf.Map[k]) == 1 {
			continue
		}
		sort.Slice(uf.Map[k], func(i, j int) bool { return uf.Map[k][i].Name > uf.Map[k][j].Name })
	}
}

func (uf *UniqueFiles) DeleteDuplicates() error {
	for k, _ := range uf.Map {
		if len(uf.Map[k]) == 1 {
			continue
		}
		for i, _ := range uf.Map[k] {
			if i == 0 {
				continue
			}
			fmt.Println("...deleting", uf.Map[k][i].Path)
			err := os.Remove(uf.Map[k][i].Path)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUniqueFilesMap() *UniqueFiles {
	return &UniqueFiles{&sync.Mutex{}, make(map[[sha512.Size]byte][]File)}
}
