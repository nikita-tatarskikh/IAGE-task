package task3

type Request []struct {
		A string
		B string
		Key string
}


type TCPMultiplier interface {
	Multiply(request *Request) string
}

type PlainService struct {
	tcp TCPMultiplier
}

func (s *PlainService) ExecuteTask3(request *Request) string {
	return s.tcp.Multiply(request)
}

func NewService(tcp TCPMultiplier) *PlainService {
	return &PlainService{tcp: tcp}
}