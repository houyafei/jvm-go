package main

import (
	"fmt"
	"io/ioutil"
	"jvmgo/ch02/classpath"
	"strings"
)

// ./ch02 -Xjre "/usr/lib/jvm/java-8-openjdk-amd64/jre" java.lang.Object
func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("hyf-jvm-version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
	//entry := &Mylist{entry:make([]string,10)}
	//findAllFile("./",entry)
	//fmt.Println(entry.entry)
	//fmt.Println(len(entry.entry))

}

type Mylist struct {
	entry []string
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data :%v\n", classData)

}
func findAllFile(pathname string, entry *Mylist) error {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Println(pathname + fi.Name()+"/")
			findAllFile(pathname+fi.Name()+"/", entry)

		} else {
			fmt.Println(fi.Name())
			entry.entry = append(entry.entry, fi.Name())
		}
	}
	return err
}
