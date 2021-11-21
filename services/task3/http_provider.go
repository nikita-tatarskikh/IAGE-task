package task3

import (
	"encoding/json"
	"log"
	"net/http"
)


type Service interface {
	ExecuteTask3(request *Request) string
}

type httpProvider struct {
	svc Service
}

type HandlersHTTP struct {
	Find http.Handler
}

func (p *httpProvider) find(w http.ResponseWriter, r *http.Request) {
	var requestData Request
	json.NewDecoder(r.Body).Decode(&requestData)
	log.Println(requestData)
	client := p.svc.ExecuteTask3(&requestData)

	w.Header().Set("Content-Type", "text; charset=utf-8")
	_, err := w.Write([]byte(client))
	if err != nil {
		log.Println("Error while writing response", err)
	}
}

func NewHTTPProvider(svc Service) HandlersHTTP {
	provider := &httpProvider{svc}
	return HandlersHTTP{
		Find: http.HandlerFunc(provider.find),
	}
}
