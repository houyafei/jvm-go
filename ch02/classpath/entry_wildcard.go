package classpath

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)


func newWildcardEntry(path string) CompositeEntry {
	// remove *
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string,info os.FileInfo,err error) error {
		if err !=nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir,walkFn)
	findAllFile(baseDir,compositeEntry)
	return compositeEntry
}

func findAllFile(pathname string, entry []Entry) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+pathListSeparator+fi.Name())
			findAllFile(pathname +pathListSeparator+ fi.Name(),entry)
		} else {
			fmt.Println(fi.Name())
		}
	}
	return err
}

