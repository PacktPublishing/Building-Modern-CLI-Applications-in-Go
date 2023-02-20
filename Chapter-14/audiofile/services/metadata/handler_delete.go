package metadata

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (m *MetadataService) deleteHandler(res http.ResponseWriter, req *http.Request) {
	value, ok := req.URL.Query()["id"]
	if !ok || len(value[0]) < 1 {
		fmt.Println("Url Param 'id' is missing")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	id := string(value[0])
	fmt.Println("deleting audio by id: ", id)

	err := m.Storage.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, fmt.Sprintf("successfully deleted audio with id: %s", id))
}
