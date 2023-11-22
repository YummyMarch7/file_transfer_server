package main

import (
	"context"
	pb "file_transfer_server/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type FileTransferServer struct {
	pb.UnimplementedFileTransServer
}

func (s *FileTransferServer) DirList(ctx context.Context, path *pb.Path) (*pb.Item, error) {
	dirs, err := os.ReadDir(path.Path)
	if err != nil {
		return &pb.Item{IsErr: true}, err
	}
	var resultDir = pb.Item{IsFile: false, Name: path.Path, SubDir: make([]*pb.Item, 0), Files: make([]*pb.Item, 0)}
	for _, item := range dirs {
		if item.IsDir() == true {
			// TODO: recursion
			currentItem := pb.Item{IsFile: false, Name: item.Name(), SubDir: make([]*pb.Item, 0)}
			resultDir.SubDir = append(resultDir.SubDir, &currentItem)
		} else {
			currentItem := pb.Item{IsFile: true, Name: item.Name(), SubDir: make([]*pb.Item, 0)}
			resultDir.Files = append(resultDir.Files, &currentItem)
		}
	}
	return &resultDir, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("failed to listen")
		fmt.Println(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFileTransServer(grpcServer, &FileTransferServer{})
	err = grpcServer.Serve(lis)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
