package files

type File struct {
	Path, Name string
}

func NewFile(path, name string) File {
	return File{path, name}
}
