// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 72999ebd2f
// Version Date: Wed Mar 17 08:36:51 UTC 2021

// Package grpc provides a gRPC client for the User service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "uas/user"
	"uas/user/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.UserServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var informationEndpoint endpoint.Endpoint
	{
		informationEndpoint = grpctransport.NewClient(
			conn,
			"user.User",
			"Information",
			EncodeGRPCInformationRequest,
			DecodeGRPCInformationResponse,
			pb.InformationResponse{},
			clientOptions...,
		).Endpoint()
	}

	endpoints := svc.NewEndpoints()
	endpoints.InformationEndpoint = informationEndpoint

	return endpoints, nil
}

// GRPC Client Decode

// DecodeGRPCInformationResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC information reply to a user-domain information response. Primarily useful in a client.
func DecodeGRPCInformationResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.InformationResponse)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCInformationRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain information request to a gRPC information request. Primarily useful in a client.
func EncodeGRPCInformationRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.InformationRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
