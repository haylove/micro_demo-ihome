package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	house "house/proto/house"
)

type House struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *House) Call(ctx context.Context, req *house.Request, rsp *house.Response) error {
	log.Info("Received House.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *House) Stream(ctx context.Context, req *house.StreamingRequest, stream house.House_StreamStream) error {
	log.Infof("Received House.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&house.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *House) PingPong(ctx context.Context, stream house.House_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&house.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
