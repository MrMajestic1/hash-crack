# HashCrack
HashCrack is a command line tool for cracking hashes written in Go with the help of the Cobra library. Hashes are attempted to be cracked by building rainbow tables based on user-supplied wordlists.
## Supported Hash Functions
Currently MD5, SHA1, SHA256 and SHA512 are supported. Support for other hash functions will be added in the future.
## Installation
Clone this repository. Then, run the following command in the cloned folder:
```
go build <installPath/hash-crack>
```
## Running
HashCrack can be run as a regular executable.
On Linux, either `./path/to/hash-crack` or `hash-crack` if it is on your PATH. 
