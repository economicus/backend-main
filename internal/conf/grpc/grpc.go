package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"main/internal/core/pb"
	"main/internal/pkg/logger"
)

type Quant struct {
	*conf
}

func New() *Quant {
	q := Quant{}
	q.conf = newConf()
	return &q
}

func (q *Quant) Request(req *pb.QuantRequest) (*pb.QuantResult, error) {
	conn, err := q.connToGrpc()
	if err != nil {
		logger.Logger.Errorf("grpc connection failed")
		return nil, err
	}

	defer conn.Close()

	client := pb.NewQuantClient(conn)
	return client.Request(q.ctx, req)
}

func (q *Quant) connToGrpc() (*grpc.ClientConn, error) {
	return grpc.Dial(q.getDSN(), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
