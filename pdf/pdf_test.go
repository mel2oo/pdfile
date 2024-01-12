package pdf

import (
	"fmt"
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
		fmt.Printf("%s context:%s\n", path, r.Content)
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

func TestPDF2(t *testing.T) {
	r, err := Parse("../testfile/nomore-double-content.pdf", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestPDF3(t *testing.T) {
	r, err := Parse("../testfile/pdfs/196d3d833f23160777fccd2a5315f96adffd600a8b45258e30126d698e46b7b6.bin", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	fmt.Printf("%scontext:%s\n", "196d3d833f23160777fccd2a5315f96adffd600a8b45258e30126d698e46b7b6", r.Content)
}

func TestPDF4(t *testing.T) {
	r, err := Parse("../testfile/pdfs/4b672deae5c1231ea20ea70b0bf091164ef0b939e2cf4d142d31916a169e8e01.pdf.bin", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	fmt.Printf("%scontext:%s\n", "4b672deae5c1231ea20ea70b0bf091164ef0b939e2cf4d142d31916a169e8e01", r.Content)
}

func TestPDF5(t *testing.T) {
	r, err := Parse("../testfile/pdfs/a57d6176ef819f20e2dfb820965943f628203067ca9b894ac94017b2a0e80383.bin", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	fmt.Printf("%scontext:%s\n", "a57d6176ef819f20e2dfb820965943f628203067ca9b894ac94017b2a0e80383", r.Content)
}

func TestPDF6(t *testing.T) {
	r, err := Parse("../testfile/pdfs/816bb2a60d8f7ff4262e22eb44bd7578bae9e82cc96825e383b72efaf1e9a508.pdf", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	fmt.Printf("%scontext:%s\n", "816bb2a60d8f7ff4262e22eb44bd7578bae9e82cc96825e383b72efaf1e9a508", r.Content)
}

func TestPDF7(t *testing.T) {
	r, err := Parse("../testfile/pdfs/1510291951580000013_PSlQoBuBUUOKOvEsPKmDE.pdf.bin", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
	fmt.Printf("%scontext:%s\n", "1510291951580000013_PSlQoBuBUUOKOvEsPKmDE", r.Content)
}

// func TestPdf6(t *testing.T) {

// 	// 定义一个扩展ASCII码
// 	extendedASCIICode := 0xfc // 这里以ASCII码值 169 为例

// 	// 将扩展ASCII码值转换为对应的Unicode码点
// 	unicodeCodePoint := rune(extendedASCIICode)

// 	// 使用 string 类型将 Unicode 码点转换为 UTF-8 编码
// 	utf8String := string(unicodeCodePoint)

// 	// 输出转换后的 UTF-8 编码
// 	fmt.Printf("UTF-8 编码: %s\n", utf8String)

// }

func TestPDFspopuer(t *testing.T) {
	// 1.pdf 文件损坏
	filepath.Walk("../testfile/pdfs/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := ExtractInPoppler(path)

		// r, err := Parse(path, "")
		if err != nil {
			return err
		}

		fmt.Printf("%s context:%s\n", path, data)
		return nil
	})
}
