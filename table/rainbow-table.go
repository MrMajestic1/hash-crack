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
package table

import (
	"hash-crack/hashes"
	"os"
	"strings"
	"sync"
)

type OutputData struct {
	HashFunction string
	Data         []string
}

type RainbowTable struct {
	HashFunction   string
	HashToPassword []*map[string]string
}

func readFile(pathToFile string) ([]byte, error) {
	filePath := pathToFile
	if pathToFile == "" {
		filePath = "table/list.txt"
	}

	inFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return inFile, nil
}

func BuildRainbowTable(pathToInputFile string, hashFunctions []string, addSalt func(string) string, maxThreads int) ([]RainbowTable, error) {
	inFile, err := readFile(pathToInputFile)
	if err != nil {
		return nil, err
	}

	res := make([]RainbowTable, len(hashFunctions))
	for i, hf := range hashFunctions {
		res[i] = RainbowTable{hf, make([]*map[string]string, maxThreads)}
	}

	stringInp := strings.Split(string(inFile[:]), "\n")
	workloadSize := len(stringInp) / maxThreads
	leftOver := len(stringInp) % maxThreads
	var wg sync.WaitGroup
	for _, rt := range res {
		for i := 0; i < maxThreads; i++ {
			wg.Add(1)
			rt.HashToPassword[i] = &map[string]string{}
			go func(i int, resi *map[string]string) {
				defer wg.Done()
				workload := stringInp[i*workloadSize : (i+1)*workloadSize]
				for _, pw := range workload {
					hash, err := hashes.GetHash(addSalt(pw), rt.HashFunction)
					if err != nil {
						return
					}
					(*resi)[hash] = pw
				}
			}(i, rt.HashToPassword[i])
		}

		wg.Wait()
		for i := len(stringInp) - leftOver; i < len(stringInp); i++ {
			hash, err := hashes.GetHash(stringInp[i], rt.HashFunction)
			if err != nil {
				return res, err
			}
			(*rt.HashToPassword[0])[hash] = stringInp[i]
		}
	}

	return res, nil
}

func ReadRainbowTable(pathToInputFile string, hashFunctions []string, maxThreads int) ([]RainbowTable, error) {
	inFile, err := readFile(pathToInputFile)
	if err != nil {
		return nil, err
	}

	maps := make([]*map[string]string, maxThreads)
	stringInp := strings.Split(string(inFile[:]), "\n")
	workloadSize := len(stringInp) / maxThreads
	leftOver := len(stringInp) % maxThreads
	var wg sync.WaitGroup
	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		maps[i] = &map[string]string{}
		go func(i int, resi *map[string]string) {
			defer wg.Done()
			workload := stringInp[i*workloadSize : (i+1)*workloadSize]
			for _, hashPw := range workload {
				pair := strings.Split(hashPw, ":")
				if len(pair) < 2 {
					continue
				}
				(*resi)[pair[0]] = pair[1]
			}
		}(i, maps[i])
	}

	wg.Wait()
	for i := len(stringInp) - leftOver; i < len(stringInp); i++ {
		pair := strings.Split(stringInp[i], ":")
		if len(pair) < 2 {
			continue
		}
		(*maps[0])[pair[0]] = pair[1]
	}

	res := make([]RainbowTable, len(hashFunctions))
	for i := 0; i < len(res); i++ {
		res[i] = RainbowTable{hashFunctions[i], maps}
	}
	return res, nil
}

func BuildRainbowTableForOutput(pathToInputFile string, hashFunctions []string, addSalt func(string) string, maxThreads int) ([]OutputData, error) {
	inFile, err := readFile(pathToInputFile)
	if err != nil {
		return nil, err
	}

	res := make([]OutputData, len(hashFunctions))
	stringInp := strings.Split(string(inFile[:]), "\n")
	inputLen := len(stringInp)
	outStrings := make([]string, len(hashFunctions)*inputLen)
	workloadSize := inputLen / maxThreads
	leftOver := inputLen % maxThreads
	var wg sync.WaitGroup
	for k, hf := range hashFunctions {
		for i := 0; i < maxThreads; i++ {
			wg.Add(1)
			go func(i int, k int) {
				defer wg.Done()
				start := i * workloadSize
				end := start + workloadSize
				outStart := k*inputLen + start
				workload := stringInp[start:end]
				for j := 0; j < len(workload); j++ {
					hash, err := hashes.GetHash(addSalt(workload[j]), hf)
					if err != nil {
						return
					}
					outStrings[outStart+j] = hash + ":" + workload[j]
				}
			}(i, k)
		}

		wg.Wait()
		for i := inputLen - leftOver; i < inputLen; i++ {
			hash, err := hashes.GetHash(outStrings[i], hf)
			if err != nil {
				return nil, err
			}
			outStrings[i] = hash + ":" + stringInp[i]
		}

		res[k] = OutputData{hf, outStrings[k*inputLen : (k+1)*inputLen]}
	}

	return res, nil
}

func (rt RainbowTable) ToString() string {
	res := rt.HashFunction + "\n"
	for i := 0; i < len(rt.HashToPassword); i++ {
		for hash, value := range *(rt.HashToPassword[i]) {
			res += hash + ":" + value + "\n"
		}
	}

	return res
}
