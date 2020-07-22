package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func (self *Classpath) parseBootAndExtClasspath(jreOptions string) {
	jreDir := getJreDir(jreOptions)
	// /usr/lib/jvm/java-8-openjdk-amd64/jre"
	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	//jre/lib/*
	jreExPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExPath)
}

func (self *Classpath) parseUserClasspath(cpOptions string) {
	if cpOptions == "" {
		cpOptions = "."
	}
	self.userClasspath = newEntry(cpOptions)
}

func getJreDir(jreOptions string) string {
	if jreOptions != "" && exists(jreOptions) {
		return jreOptions
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Parse(jreOptions, cpOptions string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOptions)
	cp.parseUserClasspath(cpOptions)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
