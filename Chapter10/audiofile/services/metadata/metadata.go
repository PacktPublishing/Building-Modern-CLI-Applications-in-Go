package metadata

import (
	"github.com/marianina8/audiofile/internal/interfaces"
	"github.com/marianina8/audiofile/storage"

	"fmt"
	"net/http"
)

type MetadataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

func CreateMetadataService(port int, storage interfaces.Storage) *MetadataService {
	mux := http.NewServeMux()
	metadataService := &MetadataService{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: mux,
		},
		Storage: storage,
	}
	mux.HandleFunc("/upload", metadataService.uploadHandler)
	mux.HandleFunc("/request", metadataService.getByIDHandler)
	mux.HandleFunc("/list", metadataService.listHandler)
	mux.HandleFunc("/delete", metadataService.deleteHandler)
	mux.HandleFunc("/search", metadataService.searchHandler)
	return metadataService
}

func Run(port int) *http.Server {
	flatfileStorage := storage.FlatFile{}
	service := CreateMetadataService(port, flatfileStorage)
	err := service.Server.ListenAndServe()
	if err != nil {
		fmt.Println("error starting api: ", err)
	}
	return service.Server
}
