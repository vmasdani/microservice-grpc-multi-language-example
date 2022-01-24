package main

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	UnimplementedGoServiceServer
}

type AppState struct {
	Counter uint
}

var appState AppState

func (s *server) Increment(ctx context.Context, in *EmptyMessage) (*MasterMessage, error) {
	appState.Counter += 1
	fmt.Println("[CALLED GoServiceServer Increment] counter:", appState.Counter)

	mMessage := MasterMessage{
		Payload: func() *string {
			p := "[Go message] " + strconv.Itoa(int(appState.Counter))
			return &p
		}(),
		Timestamp: func() *string {
			t := time.Now().Format(time.RFC3339)
			return &t
		}(),
	}

	return &mMessage, nil
}

func main() {
	go func() {
		for {
			// Java client
			fmt.Println("[From Go to Java server]", time.Now().Format(time.RFC3339))

			// Set up a connection to the server.
			func() {
				conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					fmt.Println("did not connect: %v", err)
					return
				}
				defer conn.Close()
				c := NewJavaServiceClient(conn)

				// Contact the server and print out its response.
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.Increment(ctx, &EmptyMessage{})
				if err != nil {
					fmt.Println("could not greet: %v", err)
					return
				}

				jsonR, _ := json.MarshalIndent(r, "", "  ")

				fmt.Println(string(jsonR))
			}()

			time.Sleep(time.Second * 1)

			fmt.Println()

			// C# client
			fmt.Println("[From Go to C# server]", time.Now().Format(time.RFC3339))
			fmt.Println("UNIMPLEMENTED")
			fmt.Println()

			time.Sleep(time.Second * 1)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50052))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
