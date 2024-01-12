package pdf

import (
	"bytes"
	"strings"
)

type Page Dictionary

func (page Page) Extract(output *Output) {
	d := Dictionary(page)

	// load fonts
	font_map := map[string]*Font{}
	resources, _ := d.GetDictionary("Resources")
	fonts, _ := resources.GetDictionary("Font")
	for font := range fonts {
		font_info, _ := fonts.GetDictionary(font)
		font_map[font] = NewFont(font_info)
	}

	// get contents
	if contents, ok := d.GetStream("Contents"); ok {
		page.extract(output, font_map, contents)
	} else if contents_array, ok := d.GetArray("Contents"); ok {
		for i := range contents_array {
			if contents, ok := contents_array.GetStream(i); ok {
				page.extract(output, font_map, contents)
			}
		}
	}
}

func (page Page) extract(output *Output, font_map map[string]*Font, contents []byte) {
	// create parser for parsing contents
	// fmt.Printf(string(contents))
	page_parser := NewParser(bytes.NewReader(contents), nil)

	for {
		// read next command
		command, _, err := page_parser.ReadCommand()
		if err == ErrReadError {
			break
		}

		// start of text block
		if command == KEYWORD_TEXT {
			// initial font is none
			current_font := FontDefault

			for {
				command, operands, err := page_parser.ReadCommand()
				// stop if end of stream or end of text block
				if err == ErrReadError || command == KEYWORD_TEXT_END {
					break
				}

				// handle font changes
				if command == KEYWORD_TEXT_FONT {
					font_name, _ := operands.GetName(len(operands) - 2)
					if font, ok := font_map[font_name]; ok {
						current_font = font
					} else {
						current_font = FontDefault
					}
				} else if command == KEYWORD_TEXT_SHOW_1 {

					s, _ := operands.GetString(len(operands) - 1)

					output.Content += current_font.Decode([]byte(s))

				} else if command == KEYWORD_TEXT_SHOW_2 || command == KEYWORD_TEXT_SHOW_3 {
					// decode text with current font font

					output.Content += "\n"
					s, _ := operands.GetString(len(operands) - 1)
					output.Content += current_font.Decode([]byte(s))
				} else if command == KEYWORD_TEXT_POSITION {
					// decode positioned text with current font
					var sb strings.Builder
					a, _ := operands.GetArray(len(operands) - 1)
					for i := 0; i < len(a); i += 2 {
						s, _ := a.GetString(i)

						sb.WriteString(string(s))

					}
					output.Content += current_font.Decode([]byte(sb.String())) + "\n"
				}
			}
		}
	}
}
