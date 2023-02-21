package utils

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/marianina8/audiofile/models"
)

func Print(b []byte, jsonFormat bool) ([]byte, error) {
	var err error
	if jsonFormat {
		if IsaTTY() {
			err = Pager(string(b))
			if err != nil {
				return b, fmt.Errorf("\n  paging: %v\n  ", err)
			}
		} else {
			return b, fmt.Errorf("not a tty")
		}
	} else {
		var audios models.AudioList
		var audio models.Audio
		err := json.Unmarshal(b, &audios)
		if err == nil {
			tableData, err := audios.Table()
			if err != nil {
				return b, fmt.Errorf("\n  printing table: %v\n  ", err)
			}

			if IsaTTY() && runtime.GOOS != "windows" {
				err = Pager(tableData)
				if err != nil {
					return b, fmt.Errorf("\n  paging: %v\n  ", err)
				}

			} else {
				return b, fmt.Errorf("not a tty")
			}
		} else {
			err := json.Unmarshal(b, &audio)
			if err == nil {
				tableData, err := audio.Table()
				if err != nil {
					return b, fmt.Errorf("\n  printing table: %v\n  ", err)
				}
				if IsaTTY() && runtime.GOOS != "windows" {
					err = Pager(tableData)
					if err != nil {
						return b, fmt.Errorf("\n  paging: %v\n  ", err)
					}
				} else {
					return b, fmt.Errorf("not a tty")
				}
			} else {
				return b, fmt.Errorf("unmarshalling error: %v", err)
			}
		}
	}
	return b, nil
}
