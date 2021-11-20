package task1

import (
	"encoding/json"
	"net/http"
)

type Service interface {
	ExecuteTask1(request *Request) *Response
}

type httpProvider struct {
	svc Service
}

type HandlersHTTP struct {
	Find http.Handler
}

func (p *httpProvider) find(w http.ResponseWriter, r *http.Request) {
	var requestData Request
	_ = json.NewDecoder(r.Body).Decode(&requestData)
	client := p.svc.ExecuteTask1(&requestData)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(client)
}

func NewHTTPProvider(svc Service) HandlersHTTP {
	provider := &httpProvider{svc}
	return HandlersHTTP{
		Find: http.HandlerFunc(provider.find),
	}
}