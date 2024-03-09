/*
Copyright © 2024 Quetzalcoatl

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"hash-crack/hashes"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var (
	version    string = "0.1.0"
	functions  string
	maxThreads int
	rootCmd    = &cobra.Command{
		Use:     "hashcrack",
		Short:   "HashCrack. Built by Quetzalcoatl with Cobra.",
		Long:    `HashCrack is a command line tool to crack and determine hash functions. Built by Quetzalcoatl with Cobra. Full documentation at https://github.com/MrMajestic1/hash-crack`,
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func Execute() {
	fmt.Println(rootCmd.Long)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")
	rootCmd.PersistentFlags().StringVarP(&functions, "functions", "f", strings.Join([]string{hashes.MD5, hashes.SHA1, hashes.SHA256, hashes.SHA512}, " "), "list of hash functions to try out")
	num := runtime.NumCPU()
	rootCmd.PersistentFlags().IntVar(&maxThreads, "max-threads", num, "maximum number of threads allowed")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
