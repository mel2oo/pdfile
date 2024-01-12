package pdf

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

var decoder *encoding.Decoder

func init() {
	decoder = unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
}

var FontDefault *Font = &Font{nil, map[int]string{}, 1}

type Font struct {
	Table *[256]rune
	Cmap  map[int]string
	Width int
}

func NewFont(d Dictionary) *Font {
	//  create new font object
	font := &Font{nil, map[int]string{}, 1}
	if enc, has := d["Encoding"]; has {
		if Name, ok := enc.(Name); ok {
			enc_str := string(Name)
			switch enc_str {
			case "WinAnsiEncoding":
				font.Table = &winAnsiEncoding
				return font
			case "MacRomanEncoding":
				font.Table = &macRomanEncoding
				return font
			}
		}
	}
	cmap_string, has := d.GetStream("ToUnicode")
	if !has || string(cmap_string) == "" {
		font.Table = &pdfDocEncoding
	}
	cmap := []byte(cmap_string)

	// create parser for parsing cmap
	parser := NewParser(bytes.NewReader(cmap), nil)

	for {
		// read next command
		command, operands, err := parser.ReadCommand()
		if err == ErrReadError {
			break
		}

		if command == KEYWORD_BEGIN_BF_RANGE {
			count, _ := operands.GetInt(len(operands) - 1)
			for i := 0; i < count; i++ {
				start_b := parser.ReadHexString(noDecryptor)
				if start_b == "" {
					break
				}
				font.Width = len([]byte(start_b))
				start := BytesToInt([]byte(start_b))

				end_b := parser.ReadHexString(noDecryptor)
				if end_b == "" {
					break
				}
				end := BytesToInt([]byte(end_b))
				current_offset := parser.CurrentOffset()

				value := parser.ReadHexString(noDecryptor)
				if value == "" {
					parser.Seek(current_offset, io.SeekStart)
					value_array := parser.ReadArray(noDecryptor)
					_ = value_array
					for i, j := 0, start; j <= end && i < len(value_array); i, j = i+1, j+1 {
						obj, ok := value_array[i].(String)
						if ok {
							font.Cmap[j] = obj.ToUtf8()
						}

					}
					break
				} else {
					uts8, err := decoder.Bytes([]byte(value))
					if err != nil {
						break
					}

					for j := start; j <= end; j++ {
						font.Cmap[j] = string(uts8)
					}
				}

			}
		} else if command == KEYWORD_BEGIN_BF_CHAR {
			count, _ := operands.GetInt(len(operands) - 1)
			for i := 0; i < count; i++ {
				key_b := parser.ReadHexString(noDecryptor)
				if key_b == "" {
					break
				}
				font.Width = len([]byte(key_b))
				// fmt.Println([]byte(key_b))
				key := BytesToInt([]byte(key_b))

				value := parser.ReadHexString(noDecryptor)

				font.Cmap[key] = value.ToUtf8()

			}
		}
	}

	return font
}

func (font *Font) Decode(b []byte) string {
	var s strings.Builder
	if font.Table != nil {
		r := make([]rune, 0, len(b))
		for i := 0; i < len(b); i++ {
			r = append(r, font.Table[b[i]])
		}
		return string(r)
	}

	for i := 0; i+font.Width <= len(b); i += font.Width {
		bs := b[i : i+font.Width]
		k := BytesToInt(bs)
		if v, ok := font.Cmap[k]; ok {
			s.WriteString(v)
		} else {
			s.WriteString(string(bs))
		}
	}
	return s.String()
}
