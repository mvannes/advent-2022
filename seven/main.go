package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type File struct {
	Size int
	Name string
}

type Dir struct {
	Parent *Dir
	Files  []*File
	Dirs   []*Dir
	Name   string
}

func (d *Dir) cd(name string) *Dir {
	if name == ".." {
		return d.Parent
	}
	for _, c := range d.Dirs {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (d *Dir) AddDir(dir *Dir) {
	for _, c := range d.Dirs {
		if c.Name == dir.Name {
			return
		}
	}
	d.Dirs = append(d.Dirs, dir)
}

func (d *Dir) AddFile(f *File) {
	for _, c := range d.Files {
		if c.Name == f.Name {
			return
		}
	}
	d.Files = append(d.Files, f)
}

func (d Dir) size(withSubDirs bool) int {
	var totalSize int
	for _, f := range d.Files {
		totalSize += f.Size
	}
	if !withSubDirs {
		return totalSize
	}

	for _, c := range d.Dirs {
		totalSize += c.size(true)
	}
	return totalSize
}
func (d *Dir) flattenedSubDirs() []*Dir {
	subdirFlat := []*Dir{d}
	for _, c := range d.Dirs {
		subdirFlat = append(subdirFlat, c.flattenedSubDirs()...)
	}
	return subdirFlat
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(f)

	cdRex := regexp.MustCompile("\\$ cd ([a-z./]+)$")
	dirRex := regexp.MustCompile("dir ([a-z]+)$")
	fileRex := regexp.MustCompile("(\\d+) ([a-z./]+)$")

	var initialDirName string
	// pop one, to fetch init dir.
	for scan.Scan() {
		initialDirLine := scan.Text()
		initialDirName = cdRex.FindStringSubmatch(initialDirLine)[1]
		break
	}

	mainDir := &Dir{Name: initialDirName, Parent: nil, Dirs: []*Dir{}, Files: []*File{}}
	currentDir := mainDir
	for scan.Scan() {
		line := scan.Text()
		if cdRes := cdRex.FindStringSubmatch(line); len(cdRes) > 0 {
			dirName := cdRes[1]
			currentDir = currentDir.cd(dirName)
			continue
		}
		if dirRes := dirRex.FindStringSubmatch(line); len(dirRes) > 0 {
			currentDir.AddDir(&Dir{Name: dirRes[1], Parent: currentDir, Dirs: []*Dir{}, Files: []*File{}})
			continue
		}
		if fileRes := fileRex.FindStringSubmatch(line); len(fileRes) > 0 {
			size, err := strconv.ParseInt(fileRes[1], 10, 0)
			if nil != err {
				log.Fatal(err)
			}
			currentDir.AddFile(&File{Name: fileRes[2], Size: int(size)})
			continue
		}
	}

	var sumSizes int
	for _, d := range mainDir.flattenedSubDirs() {
		if dSize := d.size(true); dSize <= 100000 {
			sumSizes += dSize
		}
	}
	fmt.Println(sumSizes)
}
