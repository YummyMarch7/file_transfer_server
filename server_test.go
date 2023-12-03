package main

import (
	pb "file_transfer_server/pb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
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

	f, e := os.Create("testfile.bin")
	if e != nil {
		fmt.Println(e)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	for _, v := range dirList.Files {
		if v.Name == "passwd" {
			stream, err := client.FileDownload(context.Background(), v)
			if err != nil {
				fmt.Println(err)
			}
			for {
				fileBlock, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
				}
				_, err = f.Write(fileBlock.BlockData)
				if err != nil {
					return
				}
				log.Println(string(fileBlock.BlockData))
			}
		}
	}
}
