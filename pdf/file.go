package pdf

import (
	"fmt"
)

type File Dictionary

func (file File) Extract(output *Output, isCommand bool) {
	d := Dictionary(file)

	// file specification can be a url or file
	fs, _ := d.GetString("FS")
	if fs == "URL" {
		if f, ok := d.GetString("F"); ok {
			output.AddUrl(f)
		}
	} else if ef, ok := d.GetDictionary("EF"); ok {
		// get the file data
		file_data, ok := ef.GetStream("F")
		if !ok {
			return
		}

		// get the file path
		f, _ := d.GetString("F")

		// dump file
		output.AddFile(f, file_data)
	} else if p, ok := d.GetString("P"); ok {
		if f, ok := d.GetString("F"); ok {
			// fmt.Fprintf(output.Files, "%s:%s\n", unknownHash, f)
			output.Command = append(output.Command,
				fmt.Sprintf("%s %s", f, p))
		}
	} else if f, ok := d.GetString("F"); ok {
		if isCommand {
			output.Command = append(output.Command,
				fmt.Sprintf("%s %s", f, p))
		}
		// fmt.Fprintf(output.Files, "%s:%s\n", unknownHash, f)
	}
}
