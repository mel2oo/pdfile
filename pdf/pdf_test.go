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
	r, err := Parse("../testfile/pdfs/7cdce94b52e431b8c5ed92d16501165e.bin", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
