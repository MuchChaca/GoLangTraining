package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var path string // path to purge
var days int    // age, in days, to purge
var test bool   // test mode toggle

func init() {
	// Set up command-line flags.
	//* Set a flag:
	// //* flag.TypeVar(pointerToVarStockingValue, "nameOfFlag", defaultValue, "description")
	flag.StringVar(&path, "p", "", "path to purge")
	flag.IntVar(&days, "d", -1, "number of days")
	flag.BoolVar(&test, "t", false, "test mode -- doesn't actually delete and log to stdout")
	//* After all flags are defined, call
	flag.Parse()

	// Tag validation
	if path == "" {
		log.Fatal("No path passed.")
	}
	if !isDir(path) {
		log.Fatal(path, "is not a valid path.")
	}
	/* if days < 1 {
		log.Fatal("Number of days must be positive, non-zero number. This is for your own protection.")
	} */
	// Set up login ??
	if test {
		log.Println("Test mode. Will not delete anything.")
		return
	}

	logFile := fmt.Sprintf("purge_%s.log", time.Now().Format("2006-01-02_150405"))
	logPath := filepath.Join(os.TempDir(), logFile) //?
	writer, err := os.Create(logPath)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(writer)
}

// isDir accepts a string (file path) and returns
// a boolean which indicates if the path is
// a valid directory.
func isDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	// Returns a boolean telling if it's a directory
	return stat.IsDir()
}

// walker implements filepath.WalkFunc
func walker(path string, info os.FileInfo, err error) error {
	cutOff, err := time.ParseDuration(fmt.Sprintf("%dh", days*24))
	if err != nil {
		log.Fatal(err)
	}
	// info.ModTime() returns a duration
	// so if it's too new, we get out
	if time.Now().Sub(info.ModTime()) <= cutOff {
		return nil
	}

	// Doesn't delete a directory unless it's empty
	if info.IsDir() {
		files, err := ioutil.ReadDir(path) // returns a slice of all the files in that directory
		if err != nil {
			log.Fatal(err)
		}

		if len(files) > 0 {
			log.Printf("directory %s contains files; skipping\n", path)
			return nil
		}
		log.Printf("directory %s is empty\n", path)
	}
	log.Printf("deleting %s\n", path)

	// If not in test mode we actually delete it
	if !test {
		err := os.Remove(path)
		if err != nil {
			log.Println("Can't delete -- ", err)
		}
	}
	return nil
}

func main() {
	filepath.Walk(path, walker)
}
