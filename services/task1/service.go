package task1

type Request struct {
	Key string `json:"key"`
	Val int `json:"val"`
}

func (req *Request) GetKey() string {
	return req.Key
}

func (req *Request) GetValue() int {
	return req.Val
}

type Response struct {
	Test int `json:"test"`
}

func (resp *Response) SetValue(val int) {
	resp.Test = val
}

type Repository interface {
	IncrementByKeyAndReturnValue(request *Request) *Response
}

type PlainService struct {
	repo Repository
}

func (s *PlainService) ExecuteTask1(request *Request) *Response {
	return s.repo.IncrementByKeyAndReturnValue(request)
}

func NewService(repo Repository) *PlainService {
	return &PlainService{repo: repo}
}