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
	"hash-crack/table"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// buildtableCmd represents the buildtable command
var (
	pathToOutputFile string
	buildtableCmd    = &cobra.Command{
		Use:   "buildtable [FLAGS] FILE",
		Short: "Build a rainbow table from provided file.",
		Long: `Build a rainbow table from provided file. The file must have the following format:

  [STRING\n...]

The output file will have the following format:

  [HASH:STRING\n...]

Flags:
	--prepend-salt=SALT  Specify a salt value to prepend during hash computations
	--append-salt=SALT   Specify a salt value to append during hash computations
	--out (-o)           Path to output file. If multiple hash functions are specified, a separate output file will be created for each hash function`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("No input specified for building rainbow table")
			}

			addSalt := func(inp string) string {
				res := inp
				if appendSalt != "" {
					res += appendSalt
				}
				if prependSalt != "" {
					res = prependSalt + res
				}

				return res
			}
			outputData, err := table.BuildRainbowTableForOutput(args[0], strings.Split(functions, " "), addSalt, maxThreads)
			if err != nil {
				fmt.Println("Error building rainbow table:", err.Error())
				os.Exit(1)
			}

			for i := 0; i < len(outputData); i++ {
				outStr := strings.Join(outputData[i].Data, "")
				dirs := strings.Split(pathToOutputFile, "/")
				fileNameSplit := strings.Split(dirs[len(dirs)-1], ".")
				fileNameSplit[0] += "_" + outputData[i].HashFunction
				dirs[len(dirs)-1] = strings.Join(fileNameSplit, ".")
				file, err := os.Create(strings.Join(dirs, "/"))
				if err != nil {
					fmt.Println("Error writing rainbow table to output file:", err.Error())
					os.Exit(1)
				}

				defer file.Close()
				_, err = file.WriteString(outStr)
				if err != nil {
					fmt.Println("Error writing rainbow table to output file:", err.Error())
					os.Exit(1)
				}
			}

			fmt.Println("Finished building rainbow table(s)")
		},
	}
)

func init() {
	buildtableCmd.Flags().StringVarP(&pathToOutputFile, "out", "o", "out.txt", "path to output file")
	buildtableCmd.Flags().StringVar(&prependSalt, "prepend-salt", "", "specify a salt value to prepend during hash computations")
	buildtableCmd.Flags().StringVar(&appendSalt, "append-salt", "", "specify a salt value to append during hash computations")
	rootCmd.AddCommand(buildtableCmd)
}
