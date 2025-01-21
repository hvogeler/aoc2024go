package warehouse

import "os"

func ReadData(file string) string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	data := string(bytes)
	return data
}