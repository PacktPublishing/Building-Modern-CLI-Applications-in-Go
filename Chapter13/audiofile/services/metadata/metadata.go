package metadata

import (
	"github.com/marianina8/audiofile/internal/interfaces"
	"github.com/marianina8/audiofile/storage"

	"fmt"
	"net/http"
	"net/http/pprof"
)

var (
	profile = false
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

	if profile {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/{action}", pprof.Index)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
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
