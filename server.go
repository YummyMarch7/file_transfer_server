package main

import (
	"context"
	"errors"
	pb "file_transfer_server/pb"
	"fmt"
	"google.golang.org/grpc"
	"io"
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
			currentItem := pb.Item{
				IsFile:   true,
				IsErr:    false,
				Name:     item.Name(),
				Size:     0,
				FullPath: path.Path + "/" + item.Name(),
				SubDir:   nil,
				Files:    nil,
			}
			resultDir.Files = append(resultDir.Files, &currentItem)
		}
	}
	return &resultDir, nil
}

func (s *FileTransferServer) FileDownload(req *pb.Item, stream pb.FileTrans_FileDownloadServer) error {
	if req.IsFile == true {
		f, err := os.Open(req.FullPath)
		if err != nil {
			return err
		}
		block := make([]byte, 1024*1024)
		pos := 0
		for {
			n, err := f.Read(block)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return err
				}
			}
			var checksum uint8 = 0
			for index := 0; index < n; index++ {
				checksum += block[index]

			}

			currentPos := pb.FileBlockPosition{
				OffsetStart: uint32(pos),
				Length:      uint32(n),
			}

			currentBlock := pb.FileBlockRespond{
				Name:      req.Name,
				FullPath:  req.FullPath,
				Position:  &currentPos,
				BlockData: block[:n],
				Checksum:  uint32(checksum),
			}
			err = stream.Send(&currentBlock)
			if err != nil {
				return err
			}
		}

		return nil
	} else {
		return errors.New("such function is not implemented")
	}
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
