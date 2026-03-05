## Huffman-encode
A simple Huffman compression library for Go with no external dependencies.

Built around io.Reader and io.Writer for maximum portability

## Instalation
```bash
go get github.com/BlueGob/huffman-enconding
```
## API reference
| Function      | Signature                                          |
|---------------|----------------------------------------------------|
| HuffmanEncode | (r io.Reader) (\*Node, error)                      |
| Compress      | (tree \*Node Cell, r io.Reader, w io.Reader) error |
| Decompress    | (r io.Reader, w io.Writer) error                   |
## Usage examples
### Compression
#### Example 1
Read data from file, compress and write to file
```go
package main

import (
	"os"
	"github.com/BlueGob/huffman-enconding"
)
func main() {
	// Open source file
	input, _ := os.Open("../../testdata/t8.shakespeare.txt")
	defer input.Close()

	// 1. Build the tree
	tree, _ := huffman.HuffmanEncode(input)

	// 2. Reset file pointer to beginning
	input.Seek(0, 0)

	// 3. Create destination file
	output, _ := os.Create("document.huff")
	defer output.Close()

	// 4. Compress
	huffman.Compress(tree, input, output)
}
```
#### Example 2
Compress and decompress without passing by files
```go
package main

import (
	"bytes"
	"fmt"
	"strings"
	"github.com/BlueGob/huffman-enconding"
)

func main() {
	data := "the quick brown fox jumps over the lazy dog"
	reader := strings.NewReader(data)

	// Build tree and dictionary
	tree, _ := huffman.HuffmanEncode(reader)
	
	// Compress into a buffer
	var compressed bytes.Buffer
	reader.Seek(0, 0) // Reset strings.Reader
	huffman.Compress(tree, reader, &compressed)

	fmt.Printf("Compressed size: %d bytes\n", compressed.Len())
}
```
### Decompression
#### Example 
```go
package main

import (
	"bytes"
	"fmt"
	"strings"
	"github.com/BlueGob/huffman-enconding"
)

func main() {
    compressedFile, _ := os.Open("compressed.bin")
    defer compressedFile.Close()

    outputFile, _ := os.Create("restored.txt")
    defer outputFile.Close()

    huffman.Decompress(compressedFile, outputFile)
}
```
