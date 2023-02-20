package metadata

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (m *MetadataService) listHandler(res http.ResponseWriter, req *http.Request) {
	audioFiles, err := m.Storage.List()
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
