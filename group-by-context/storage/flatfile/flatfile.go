package flatfile

import (
	"audiofile/models"
	"fmt"
)

func Upload(bytes []byte, tag string) (int, error) {
	// generate guid
	// copy file to configured storage path by tag name or id
	// extract metadata (go function that )
	// save metadata alongside file in storage
	return 0, nil
}

func List() ([]models.Audio, error) {
	fmt.Println("Listing")
	return nil, nil
}

func Delete(id string, tag string) error {
	fmt.Println("Deleting")
	return nil
}
