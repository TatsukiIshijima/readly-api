package server

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
	"testing"
)

func NewTestMetadata(
	grpcGatewayUserAgentHeaderValue string,
	userAgentHeaderValue string,
	xForwardedHeaderValue string,
) metadata.MD {
	md := map[string]string{}
	if len(grpcGatewayUserAgentHeaderValue) != 0 {
		md[grpcGatewayUserAgentHeader] = grpcGatewayUserAgentHeaderValue
	}
	if len(userAgentHeaderValue) != 0 {
		md[userAgentHeader] = userAgentHeaderValue
	}
	if len(xForwardedHeaderValue) != 0 {
		md[xForwardedForHeader] = xForwardedHeaderValue
	}
	return metadata.New(md)
}

func TestNewMetadataFrom(t *testing.T) {
	testCases := []struct {
		name  string
		input metadata.MD
		want  *Metadata
	}{
		{
			name: "success when access from http",
			input: NewTestMetadata(
				"",
				"PostmanRuntime/7.43.0",
				"127.0.0.1"),
			want: &Metadata{
				UserAgent: "PostmanRuntime/7.43.0",
				IPAddress: "127.0.0.1",
			},
		},
		{
			name: "success when access from grpc",
			input: NewTestMetadata(
				"grpc-node-js/1.11.0-postman.1",
				"",
				"127.0.0.1"),
			want: &Metadata{
				UserAgent: "grpc-node-js/1.11.0-postman.1",
				IPAddress: "127.0.0.1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tc.input)
			got := newMetadataFrom(ctx)
			require.Equal(t, tc.want, got)
		})
	}
}
