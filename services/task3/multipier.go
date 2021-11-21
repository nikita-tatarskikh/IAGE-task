package task3

import (
	"bufio"
	"fmt"
	"log"
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
		log.Println(forTCP)
		conn, _ := net.Dial("tcp", "127.0.0.1:8081")
		fmt.Fprintf(conn, forTCP)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
		conn.Close()
	}
	return "forTCP"
}


func NewMultiply() *Multiply {
	return &Multiply{}
}

