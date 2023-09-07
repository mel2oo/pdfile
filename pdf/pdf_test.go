package pdf

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPDFs(t *testing.T) {
	// 1.pdf 文件损坏
	filepath.Walk("../testfile/pdfs/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		r, err := Parse(path, "")
		if err != nil {
			return err
		}
		t.Log(r)
		return nil
	})
}

func TestPDF(t *testing.T) {
	r, err := Parse("../testfile/pdfs/816bb2a60d8f7ff4262e22eb44bd7578bae9e82cc96825e383b72efaf1e9a508.pdf", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
