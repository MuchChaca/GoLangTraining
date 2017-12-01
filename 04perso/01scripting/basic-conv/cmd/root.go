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
	"fmt"
	"log"
	"os"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var source = make([]string, 2)
var dest string
var typeConv string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "basic-conv",
	Short: "Basic unit converter",
	Long: `***** Basic converter *****
For example:

	bconv -s [value source],[unit source] -d [unit destination] -t [type conversion]

	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//* maybe call some func there
		// fmt.Println(len(source))
		convert(source, dest)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//* Usages
	// usageType := ``
	// usageUnit := ``

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.basic-conv.yaml)")

	// MyAdds
	RootCmd.Flags().StringSliceVarP(&source, "source", "s", []string{"", ""}, "The value and unit to be converted -- try the help command '-h' to learn more")
	RootCmd.Flags().StringVarP(&dest, "destination", "d", "", "The unit to convert to -- see the list of units supported")
	RootCmd.Flags().StringVarP(&typeConv, "type", "t", "", "The type of conversion to operate (e.g. mass) -- use -h to see all supported types.")

	// fmt.Println(RootCmd.Flags())

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "o", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".basic-conv" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".basic-conv")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//* convert returns the value converted plus the unit converted to
//* will return the result and the 'dest string'
func convert(source []string, dest string) error {
	//	// source[0] needs to be the value source
	//	// source[1] needs to be the unit source
	//	// dest is the unit destination
	switch {
	case len(source) != 2:
		log.Fatalln("ERROR - The number of arguments to be converted is incorrect.")
	case source[0] == "":
		log.Fatalln("ERROR - The value to be converted need to be provided in -s.")
	case source[1] == "":
		log.Fatalln("ERROR - The unit to be converted need to be provided in -s.")
	}

	// fmt.Println(source)

	// var result float64
	var converted float64

	//* Handle the differents units
	// TODO: complete
	switch typeConv {
	case "mass":
		converted = convertMass(source[0], source[1], dest)
	}
	valueFinal := strconv.FormatFloat(converted, 'f', -1, 64)

	fmt.Println(`Conversion :
` + source[0] + source[1] + ` = ` + valueFinal + `
		`)

	return nil
}

//* ---- Converters ---- *//
//	// Check the type of conversion and makes one calculation for a type
//	// bring back the source value to the default unit, then in the end
//	// converts from this default unit to the desired one

// convertMass converts all mass related units
func convertMass(valueS string, unitS string, unitD string) float64 {
	value, errConv := strconv.ParseFloat(valueS, 64)
	if errConv != nil {
		log.Fatal("Unable to process the value provided.")
	}

	// valueS string -> value float64

	// bring it to grammes in the switch
	switch unitS {
	case "kg":
		value = value / 0.001
	}

	// final calcul
	valueConv := 1.0
	// according to the unit dest
	switch dest {
	case "hg": //hectograms
		valueConv = 0.01
	case "lb": // pounds
		valueConv = 0.00220462
	}

	return value * valueConv
}
