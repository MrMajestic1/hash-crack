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
	"hash-crack/crack"
	"hash-crack/table"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// crackCmd represents the crack command
var (
	prependSalt  string
	appendSalt   string
	rainbowTable string
	wordlist     string

	crackCmd = &cobra.Command{
		Use:   "crack",
		Short: "Crack hash(es) specified as input",
		Long: `Crack hash(es) specified as input.

Usage:
	crack [FLAGS] HASH[ HASH...]
		
Flags:
	--prepend-salt=SALT  Specify a salt value to prepend during hash computations
	--append-salt=SALT   Specify a salt value to append during hash computations
	--wordlist=FILE      Specify a wordlist to build rainbow table
	--rainbow-table=FILE Specify a file containing a rainbow table`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("No arguments specified. Please specify file of hashes or hash string to crack")
				return
			}

			fmt.Println("Attempting to crack hashes:", args)
			fmt.Println("Hash functions used:", functions)
			functionsArray := strings.Split(functions, " ")

			var hashToPw []table.RainbowTable
			var err error
			var addSalt = func(inp string) string {
				res := inp
				if appendSalt != "" {
					res += appendSalt
				}
				if prependSalt != "" {
					res = prependSalt + res
				}

				return res
			}
			start := time.Now()
			defer calcRuntime(start)
			if rainbowTable != "" {
				hashToPw, err = table.ReadRainbowTable(rainbowTable, functionsArray, maxThreads)
			} else {
				hashToPw, err = table.BuildRainbowTable(wordlist, functionsArray, addSalt, maxThreads)
			}

			if err != nil {
				fmt.Println("Error while building rainbow tables:", err.Error())
				os.Exit(1)
			}

			for i := 0; i < len(args); i++ {
				inputVal, hashFunction := crack.CrackHash(args[i], hashToPw)
				fmt.Println(args[i], ":", inputVal, "Hash function:", hashFunction)
			}
		},
	}
)

func calcRuntime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("Runtime:", elapsed, "with", maxThreads, "threads")
}

func init() {
	crackCmd.Flags().StringVar(&prependSalt, "prepend-salt", "", "specify a salt value to prepend during hash computations")
	crackCmd.Flags().StringVar(&appendSalt, "append-salt", "", "specify a salt value to append during hash computations")
	crackCmd.Flags().StringVar(&wordlist, "wordlist", "", "specify a wordlist file to build rainbow table")
	crackCmd.Flags().StringVar(&rainbowTable, "rainbow-table", "", " build rainbow table from a given file")

	rootCmd.AddCommand(crackCmd)
}
