package main

import (
	"fmt"
	"os"
)


//type FileInfo interface {
//	Name() string       // base name of the file
//	Size() int64        // length in bytes for regular files; system-dependent for others
//	Mode() FileMode     // file mode bits
//	ModTime() time.Time // modification time
//	IsDir() bool        // abbreviation for Mode().IsDir()
//	Sys() interface{}   // underlying data source (can return nil)
//}

//// A FileMode represents a file's mode and permission bits.
//// The bits have the same definition on all systems, so that
//// information about files can be moved from one system
//// to another portably. Not all bits apply to all systems.
//// The only required bit is ModeDir for directories.
//type FileMode uint32

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: permissions filename\n")
		return
	}

	filename := arguments[1]
	info, _ := os.Stat(filename)
	mode := info.Mode()

	fmt.Printf( "%o\n", mode) //410000666
	fmt.Printf( "%b\n", mode) //0000 0100   0010 0000   0000 0001   1011 0110
	fmt.Println(filename, "mode is", mode.String()[1:10])


}
