package task2

type Request struct {
	S string
	Key string
}

func (req *Request) GetS() string {
	return req.S
}

func (req *Request) GetKey() string {
	return req.Key
}

type Security interface {
	CreateSign(request *Request) string
}

type PlainService struct {
	sec Security
}

func (s *PlainService) ExecuteTask2(request *Request) string {
	return s.sec.CreateSign(request)
}

func NewService(sec Security) *PlainService {
	return &PlainService{sec: sec}
}