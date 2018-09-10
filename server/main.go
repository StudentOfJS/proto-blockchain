package main

import (
	"log"
	"net"

	"github.com/studentofjs/bchain/proto"
	"github.com/studentofjs/bchain/server/blockchain"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen on port 8080 port %v", err)
	}

	srv := grpc.NewServer()

	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}

type Server struct {
	Blockchain *blockchain.Blockchain
}

func (s *Server) AddBlock(c context.Context, i *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(i.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetBlockchain(c context.Context, i *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}
