package main

import (
	pb "file_transfer_server/pb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
import "testing"

func TestMain(m *testing.M) {
	//var opts []grpc.DialOption
	conn, err := grpc.Dial("127.0.0.1:8899", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("err connect")
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewFileTransClient(conn)
	req := pb.Path{Path: "/etc", Depth: 1}
	dirList, eee := client.DirList(context.Background(), &req)
	if eee != nil {
		fmt.Println("Dir List error")
	}
	fmt.Println(dirList)
}
