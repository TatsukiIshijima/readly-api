package server

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	IPAddress string
}

func newMetadataFrom(ctx context.Context) *Metadata {
	meta := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			meta.UserAgent = userAgents[0]
		} else if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			meta.UserAgent = userAgents[0]
		}
		if ipAddresses := md.Get(xForwardedForHeader); len(ipAddresses) > 0 {
			meta.IPAddress = ipAddresses[0]
		}
	}

	if meta.IPAddress == "" {
		if ip, ok := peer.FromContext(ctx); ok {
			meta.IPAddress = ip.Addr.String()
		}
	}

	return meta
}
