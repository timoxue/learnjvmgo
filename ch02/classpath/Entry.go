package classpath

import "os"

const pathListSeparator = string(os.PathListSeparator)

//Entry as class path
type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry() {}
