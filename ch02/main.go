package main

import (
	"fmt"
	"io/ioutil"
	"os"
)
const pathListSeparator string = string(os.PathListSeparator)
func main() {
	//cmd := parseCmd()
	//if cmd.versionFlag {
	//	fmt.Println("hyf-jvm-version 0.0.1")
	//}else if cmd.helpFlag || cmd.class == "" {
	//	printUsage()
	//}else {
	//	startJVM(cmd)
	//}
	entry := &Mylist{entry:make([]string,10)}
	findAllFile("./",entry)
	fmt.Println(entry.entry)
	fmt.Println(len(entry.entry))

}

type Mylist struct {
	entry []string

}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
}
func findAllFile(pathname string, entry *Mylist) error {
	rd, err := ioutil.ReadDir(pathname)
	if err!=nil {
		return err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Println(pathname + fi.Name()+"/")
			findAllFile(pathname + fi.Name()+"/",entry)

		} else {
			fmt.Println(fi.Name())
			entry.entry = append(entry.entry,fi.Name())
		}
	}
	return err
}
