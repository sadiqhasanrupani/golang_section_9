package filemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func LoadJsonData[T any](filePath string) (T, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	var data T
	err = json.Unmarshal(file, &data)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	return data, nil
}

func WriteJsonData(path string, data any) error { // you can use any or interface{}, both are same.
	file, err := os.Create(path)

	if err != nil {
		file.Close()
		return errors.New("failed to create a file")
	}

	// now we will write json data using encoding/json package.
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("an error occurred while encoding a data into json file")
	}

	file.Close()
	fmt.Printf("\nFile has been created on %v path", path)
	return nil
}
