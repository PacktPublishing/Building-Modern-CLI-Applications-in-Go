package metadata

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (m *MetadataService) getByIDHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["id"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := string(value[0])
	fmt.Println("getByIDHandler, requesting id: ", id)

	audio, err := m.Storage.GetByID(id)
	if err != nil {
		fmt.Println("getByIDHandler, error getting by id: ", err.Error())
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no such file or directory") {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	audioString, err := audio.JSON()
	if err != nil {
		fmt.Println("getByIDHandler, error marshalling audio by id to json: ", err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("getByIDHandler, response: ", audioString)
	io.WriteString(res, audioString)
}
