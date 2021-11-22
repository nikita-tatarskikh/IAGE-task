package task3

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Multiply struct {

}

func (m *Multiply) Multiply(request *Request) string {
	for _, req := range *request {
		var s []string
		s = append(s, req.A)
		s = append(s, req.B)
		s = append(s, "\n")
		forTCP := strings.Join(s, "")
		conn, _ := net.Dial("tcp", "127.0.0.1:8081")
		fmt.Fprintf(conn, forTCP)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Task3 - Message from tcp server: " + message)
		conn.Close()
	}
	return "Пункт выполнен частично. Результат в логе."
}


func NewMultiply() *Multiply {
	return &Multiply{}
}

