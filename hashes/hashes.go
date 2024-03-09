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
package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"strings"
)

const (
	MD5    = "MD5"
	SHA1   = "SHA1"
	SHA256 = "SHA256"
	SHA512 = "SHA512"
)

func GetHash(input string, hashFunction string) (string, error) {
	var h hash.Hash
	hf := strings.ToUpper(hashFunction)

	if hf == MD5 {
		h = md5.New()
	} else if hf == SHA1 {
		h = sha1.New()
	} else if hf == SHA256 {
		h = sha256.New()
	} else if hf == SHA512 {
		h = sha512.New()
	} else {
		return "", errors.New("unsupported hash function specified")
	}

	io.WriteString(h, input)
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}
