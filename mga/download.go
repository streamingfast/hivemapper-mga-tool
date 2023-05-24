package mga

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func Download(mgaServiceURL string) (mgaData []byte, hash string, err error) {
	resp, err := http.Get(mgaServiceURL)
	if err != nil {
		err = fmt.Errorf("downloading MGA data: %w", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("reading MGA data: %w", err)
		return
	}
	h := md5.New()
	h.Write(body)
	hash = hex.EncodeToString(h.Sum(nil))

	return
}
