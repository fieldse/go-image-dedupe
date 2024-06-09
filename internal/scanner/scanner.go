// Scan directories for images

package scanner

import "time"

type File struct {
	BaseName  string
	FilePath  string
	FileSize  int64
	Created   time.Time
	Updated   time.Time
	Extension string
	AbsPath   string
}

type DirList struct {
	DirName string
	SubDirs []DirList
	Files   []File
	AbsPath string
}

func ScanForImages(imgDir string) {
}
