package file

import (
	"io"
	"os"
)

func Read(fileName string) (content string, err error) {
	if fileName == "" {
		return "", nil
	}
	myfile, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer myfile.Close()
	var cont []byte
	buf := make([]byte, 2)
	for {
		n, err := myfile.Read(buf)
		cont = append(cont, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(cont), nil
}
