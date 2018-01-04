package classpath

import (
	"os"
	"path/filepath"
)

//Classpath class path for loading class
type Classpath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

//Parse the jre path and class path
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//ReadClass read class from all class path
func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	classFullName := className + ".class"
	if data, entry, err := classpath.bootClassPath.readClass(classFullName); err == nil {
		return data, entry, err
	}
	if data, entry, err := classpath.extClassPath.readClass(classFullName); err == nil {
		return data, entry, err
	}
	return classpath.userClassPath.readClass(classFullName)
}

func (classpath *Classpath) String() string {
	return classpath.userClassPath.String()
}

func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	//jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	classpath.bootClassPath = newWildCardEntry(jreLibPath)
	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	classpath.extClassPath = newWildCardEntry(jreExtPath)
}

func (classpath *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	classpath.userClassPath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
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
