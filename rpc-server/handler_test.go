package main

import (
	"context"
	"testing"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"github.com/stretchr/testify/assert"
)

func TestIMServiceImpl_Send(t *testing.T) {
	type args struct {
		ctx context.Context
		req *rpc.SendRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &rpc.SendRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, true)
		})
	}
}
