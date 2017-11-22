// Copyright © 2017 Sebastien Bastide
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a new file",
	Long:  ``,
}

var fileName, path string

func init() {
	configCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// TODO: finish this docs on https://godoc.org/github.com/spf13/pflag#StringVarP && https://github.com/BitBalloon/bitballoon-cli/blob/master/create.go
	createCmd.Flags().String(&fileName)
}

// Create handles the creation of the file according to the existence of a path or not
// Guess it makes the default action of the command?
func Create(cmd *cobra.Command, args []string) {
	// code
	var name, path string
	switch len(args) {
	case 2:
		path = args[1]
		fallthrough
	case 1:
		name = args[0]
	default:
		log.Println("Encoutered error in create parameters")
		panic()
	}

	// WriteFile writes data to a file named by filename.
	// If the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing.
	err := ioutil.WriteFile(path + name)
	if err != nil {
		log.Fatalf("Creation of %s%s failed:\n\t%v\n", path, name, err)
	}
	log.Printf("Creation of %s%s succeeded!", path, name)
}
