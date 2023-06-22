package mga

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func Download(mgaServiceURL string) (mgaData []byte, hash string, err error) {
	fmt.Println("downloading MGA data from:", mgaServiceURL)
	// resp, err := http.Get(mgaServiceURL)
	// if err != nil {
	// 	err = fmt.Errorf("downloading MGA data: %w", err)
	// 	return
	// }

	// if resp.StatusCode != http.StatusOK {
	// 	err = fmt.Errorf("downloading MGA failed with cide: %d", resp.StatusCode)
	// 	return
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	err = fmt.Errorf("reading MGA data: %w", err)
	// 	return
	// }

	// fmt.Println("read body:", len(body))
	// mgaData = body

	// h := md5.New()
	// h.Write(body)
	// hash = hex.EncodeToString(h.Sum(nil))

	return
}
