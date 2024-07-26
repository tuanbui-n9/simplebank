package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/tuanbui-n9/simplebank/token"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context, accessibleRoles []string) (*token.Payload, error) {
	metadata, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}
	value := metadata.Get(authorizationHeader)

	if len(value) == 0 {
		return nil, fmt.Errorf("authorization token is not provided")
	}

	authHeader := value[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("authorization token is not provided")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	if !hasPermission(payload.Role, accessibleRoles) {
		return nil, fmt.Errorf("forbidden")
	}

	return payload, nil
}

func hasPermission(userRole string, accessibleRoles []string) bool {
	for _, role := range accessibleRoles {
		if userRole == role {
			return true
		}
	}

	return false
}
