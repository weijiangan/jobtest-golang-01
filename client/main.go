package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	pb "github.com/weijiangan/bruno-test/brunotest"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	if (len(os.Args) != 6 || os.Args[1] != "-s") && !(len(os.Args) == 5 && os.Args[1] == "-q") {
		fmt.Println("Usage: client [CLIENT_IP] [SERVER_IP] TAGS MESSAGE\n\n" +
			"Where TAGS is a JSON key/value pair")
		return
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		sugar.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAppClient(conn)

	// Switch base on cli args
	switch os.Args[1] {
	case "-s":
		clientIp := os.Args[2]
		serverIp := os.Args[3]
		tagStr := os.Args[4]
		msg := os.Args[5]

		var tagMap map[string]string
		json.Unmarshal([]byte(tagStr), &tagMap)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Send(ctx, &pb.AuditEvent{
			ClientIp: clientIp,
			ServerIp: serverIp,
			Tag:      tagMap,
			Message:  msg,
		})
		if err != nil {
			sugar.Fatalf("Could not send: %v", err)
		}
		sugar.Infof("StatusCode: %d Message: %s", r.StatusCode, r.Message)

	case "-q":
		clientIp := os.Args[2]
		serverIp := os.Args[3]
		tagStr := os.Args[4]

		var tagMap map[string]string
		json.Unmarshal([]byte(tagStr), &tagMap)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		stream, err := c.Query(ctx, &pb.QueryParam{
			ClientIp: clientIp,
			ServerIp: serverIp,
			Tag:      tagMap,
		})
		if err != nil {
			sugar.Fatalf("Could not send: %v", err)
		}
		for {
			auditEvent, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				sugar.Fatalf("%v.Query(_) = _, %v", c, err)
			}
			fmt.Printf("%+v\n", auditEvent)
		}
	}
}
