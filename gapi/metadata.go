package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgent = "grpcgateway-user-agent"
	userAgentHeader      = "user-agent"
	xForwardedFor        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetaData(ctx context.Context) *Metadata {
	mtdata := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := md.Get(grpcGatewayUserAgent); len(userAgents) > 0 {
			mtdata.UserAgent = userAgents[0]
		}
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdata.UserAgent = userAgents[0]
		}
		if clientIPs := md.Get(xForwardedFor); len(clientIPs) > 0 {
			mtdata.ClientIP = clientIPs[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdata.ClientIP = p.Addr.String()
	}
	return mtdata
}
