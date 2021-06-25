package main

import (
	"fmt"
	"github.com/moonetic/grpc-proto-test"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/send_message", func(writer http.ResponseWriter, request *http.Request) {

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(os.Getenv("SERVER_HOST"), grpc.WithInsecure())
		if err != nil {
			fmt.Fprintf(writer, "could not connect: %s", err)
			return
		}
		defer conn.Close()

		c := message.NewMessageServiceClient(conn)

		message := message.Message{Body: "Test message"}

		response, err := c.SendMessage(context.Background(), &message)
		if err != nil {
			fmt.Fprintf(writer, "error when calling SendMessage: %s", err)
			return
		}

		fmt.Fprintf(writer, "Response from server: %s", response.Body)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
