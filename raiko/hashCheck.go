package raiko

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
)

func HashSum(path string) error {
	h := md5.New()          // Creating hash
	f, err := os.Open(path) // Openning the banner file
	if err != nil {
		return errors.New("unable to open file")
	}
	defer f.Close()
	_, err = io.Copy(h, f) // Copying into hash
	if err != nil {
		return err
	}

	hashSum := fmt.Sprintf("%x", h.Sum(nil))                                                                                                             // getting our hashsum
	if hashSum != "a49d5fcb0d5c59b2e77674aa3ab8bbb1" && hashSum != "ac85e83127e49ec42487f272d9b9db8b" && hashSum != "86d9947457f6a41a18cb98427e314ff8" { // Comparing them with the original hash values
		return errors.New("font file is corrupted")
	}
	return nil
}
