package grpcclient

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var (
	defaultTimeout = 120 * time.Second
)

func getTimeoutContext(timeout time.Duration) (context.Context, func()) {
	return context.WithTimeout(context.Background(), timeout)
}

func getDefaultCallOptions() grpc.EmptyCallOption {
	return grpc.EmptyCallOption{}
}
