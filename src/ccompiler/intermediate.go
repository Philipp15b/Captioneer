package ccompiler

import (
	"encoding/binary"
	"fmt"
	"io"
	"unicode/utf16"
)

const captionTemplate = `"lang"
{ 
"Language" "%v" 
"Tokens" 
{
%v
}
}`

func CompileIntermediate(captions map[string]string, language string) string {
	tokens := ""
	for key, value := range captions {
		tokens += fmt.Sprintf("\t\t\"%v\"\t\t\"%v\"\n", key, value)
	}
	return fmt.Sprintf(captionTemplate, language, tokens)
}

func WriteIntermediate(w io.Writer, intermediate string) error {
	_, err := w.Write([]byte{0xFF, 0xFE}) // Write BOM (Little Endian)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, utf16.Encode([]rune(intermediate)))
	return err
}
