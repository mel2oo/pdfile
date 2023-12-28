package pdf

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/h2non/filetype"
	pdfcpuapi "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

type Output struct {
	Content    string     `json:"content"`
	Command    []string   `json:"command"`
	Javascript []string   `json:"javascript"`
	URLs       []string   `json:"urls"`
	Files      []Embedded `json:"-"`
}

type Embedded struct {
	Name string        `json:"name"`
	Type string        `json:"type"`
	Hash string        `json:"hash"`
	File io.ReadWriter `json:"-"`
}

func Parse(filepath string, password string) (*Output, error) {
	// open the pdf
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create a new parser
	output := &Output{}
	parser := NewParser(file, output)

	// load the pdf
	if err := parser.Load(password); err != nil {
		return nil, err
	}

	// extract and dump all objects
	for object_number, xref_entry := range parser.Xref {
		if xref_entry.Type == XrefTypeIndirectObject {
			object := parser.GetObject(object_number)
			object.Extract(output)
			output.AddFile("", object.Stream)
		}
	}

	// use github.com/ledongthuc/pdf
	// f, r, err := pdf.Open(filepath)
	// if err == nil {
	// 	defer f.Close()

	// 	b, err := r.GetPlainText()
	// 	if err == nil {
	// 		var buf bytes.Buffer
	// 		buf.ReadFrom(b)
	// 		output.Content = buf.String()
	// 	}
	// }

	// use github.com/pdfcpu/pdfcpu
	extractImgesWithpdfcpu(filepath, output)
	return output, nil
}

func extractImgesWithpdfcpu(inFile string, output *Output) error {
	f, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return pdfcpuapi.ExtractImages(f, nil, func(i1 model.Image, b bool, i2 int) error {
		data, err := io.ReadAll(i1)
		if err != nil {
			return err
		}
		output.AddFile("", data)
		return nil
	}, nil)
}

// chiRex is a regexp to replace chinese characters
var chiRex = regexp.MustCompile("[\u4e00-\u9fa5]")

func (o *Output) AddUrl(rawUrl string) {
	U, err := url.Parse(rawUrl)
	if err != nil {
		return
	}
	// check if the host is chinese
	if chiRex.MatchString(U.Host) {
		// remove chinese characters
		U.Host = chiRex.ReplaceAllString(U.Host, "")
		rawUrl = U.String()
	}

	for _, v := range o.URLs {
		if v == rawUrl {
			return
		}
	}
	o.URLs = append(o.URLs, rawUrl)
}

func (o *Output) AddFile(name string, data []byte) {
	if len(data) == 0 {
		return
	}

	hash := md5.New()
	hash.Write(data)
	md5sum := hex.EncodeToString(hash.Sum(nil))

	for _, f := range o.Files {
		if f.Hash == md5sum {
			return
		}
	}

	magic, err := filetype.Get(data)

	if magic.Extension == filetype.Unknown.Extension &&
		strings.Contains(strings.ToLower(string(data)), "<script>") {
		magic.Extension = "js"
		o.Javascript = append(o.Javascript, string(data))

	}

	if err != nil ||
		filetype.IsFont(data) ||
		magic.Extension == filetype.Unknown.Extension {
		return
	}

	if len(name) == 0 {
		name = fmt.Sprintf("%s.%s", md5sum, magic.Extension)
	}

	o.Files = append(o.Files, Embedded{
		Name: name,
		Type: magic.Extension,
		Hash: md5sum,
		File: bytes.NewBuffer(data),
	})
}

func (o *Output) Dump(dir string) {
	if len(dir) > 0 {
		os.RemoveAll(dir)
		os.MkdirAll(dir, 0755)
	}

	for _, emb := range o.Files {
		path := path.Join(dir, fmt.Sprintf("%s-%s", emb.Hash, emb.Name))

		fi, err := os.Create(path)
		if err != nil {
			continue
		}
		defer fi.Close()

		io.Copy(fi, emb.File)
	}
}
