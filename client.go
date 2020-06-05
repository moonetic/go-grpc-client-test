package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"moonetic.com/grpc-client-test/message"
	"net/http"
)

func main() {

	http.HandleFunc("/send_message", func(writer http.ResponseWriter, request *http.Request) {

		var conn *grpc.ClientConn

		conn, err := grpc.Dial(":9000", grpc.WithInsecure())
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
