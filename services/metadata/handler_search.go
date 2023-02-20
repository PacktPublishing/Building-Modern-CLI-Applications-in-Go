package metadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (m *MetadataService) searchHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["searchFor"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'searchFor' is missing")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	searchFor := string(value[0])
	fmt.Println("searching for audio containing the string: ", searchFor)

	audioFiles, err := m.Storage.Search(searchFor)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(audioFiles)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(jsonData), "", "    ")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(res, prettyJSON.String())
}
