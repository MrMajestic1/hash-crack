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
package crack

import "hash-crack/table"

func CrackHash(hash string, rainbowTables []table.RainbowTable) (string, string) {
	for i := 0; i < len(rainbowTables); i++ {
		rt := (rainbowTables)[i]

		for _, htp := range rt.HashToPassword {
			if (*htp)[hash] != "" {
				return (*htp)[hash], rt.HashFunction
			}
		}
	}

	return "Not able to crack hash", "No hash function"
}
