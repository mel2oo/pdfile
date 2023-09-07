package pdf

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPDFs(t *testing.T) {
	// 1.pdf 文件损坏
	//
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
