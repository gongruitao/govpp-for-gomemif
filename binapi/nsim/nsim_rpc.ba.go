// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package nsim

import (
	"context"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  nsim.
type RPCService interface {
	NsimConfigure(ctx context.Context, in *NsimConfigure) (*NsimConfigureReply, error)
	NsimCrossConnectEnableDisable(ctx context.Context, in *NsimCrossConnectEnableDisable) (*NsimCrossConnectEnableDisableReply, error)
	NsimOutputFeatureEnableDisable(ctx context.Context, in *NsimOutputFeatureEnableDisable) (*NsimOutputFeatureEnableDisableReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) NsimConfigure(ctx context.Context, in *NsimConfigure) (*NsimConfigureReply, error) {
	out := new(NsimConfigureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NsimCrossConnectEnableDisable(ctx context.Context, in *NsimCrossConnectEnableDisable) (*NsimCrossConnectEnableDisableReply, error) {
	out := new(NsimCrossConnectEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) NsimOutputFeatureEnableDisable(ctx context.Context, in *NsimOutputFeatureEnableDisable) (*NsimOutputFeatureEnableDisableReply, error) {
	out := new(NsimOutputFeatureEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
