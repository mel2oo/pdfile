package pdf

import (
	"context"
	"crypto/rand"
	"math/big"
	"os"
	"os/exec"
	"time"
)

func ExecCommandWithTimeout(timeout time.Duration, name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		if err == context.DeadlineExceeded || err == exec.ErrNotFound {
			return nil, err
		}
	}

	return out, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[index.Int64()]
	}
	return string(result)
}

// ExtractInPopplerTsv Access raw stdout content from Poppler
func ExtractInPoppler(pdfpath string) (data []byte, err error) {
	savepath := pdfpath + generateRandomString(8) + "text.data"
	params := []string{
		"-enc", "UTF-8",
		pdfpath,  // Read from stdin
		savepath, // Write to stdout
	}
	defer os.Remove(savepath)

	data, err = ExecCommandWithTimeout(time.Second*15, "pdftotext", params...)
	if err != nil {
		return data, err
	}

	return os.ReadFile(savepath)
}
