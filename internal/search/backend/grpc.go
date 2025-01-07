package backend

import (
	"context"
	"io"

	"github.com/sourcegraph/zoekt"
	proto "github.com/sourcegraph/zoekt/grpc/protos/zoekt/webserver/v1"
	"github.com/sourcegraph/zoekt/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// zoektGRPCClient is a zoekt.Streamer that uses gRPC for its RPC layer
type zoektGRPCClient struct {
	endpoint string
	client   proto.WebserverServiceClient

	// We capture the dial error to return it lazily.
	// This allows us to treat Dial as infallible, which is
	// required by the interface this is being used behind.
	dialErr error
}

var _ zoekt.Streamer = (*zoektGRPCClient)(nil)

func (z *zoektGRPCClient) StreamSearch(ctx context.Context, q query.Q, opts *zoekt.SearchOptions, sender zoekt.Sender) error {
	if z.dialErr != nil {
		return z.dialErr
	}

	req := &proto.StreamSearchRequest{
		Request: &proto.SearchRequest{
			Query: query.QToProto(q),
			Opts:  opts.ToProto(),
		},
	}

	ss, err := z.client.StreamSearch(ctx, req)
	if err != nil {
		return convertError(err)
	}

	for {
		msg, err := ss.Recv()
		if err != nil {
			return convertError(err)
		}

		var repoURLS map[string]string      // We don't use repoURLs in Khulnasoft
		var lineFragments map[string]string // We don't use lineFragments in Khulnasoft

		sender.Send(zoekt.SearchResultFromProto(msg.GetResponseChunk(), repoURLS, lineFragments))
	}
}

func (z *zoektGRPCClient) Search(ctx context.Context, q query.Q, opts *zoekt.SearchOptions) (*zoekt.SearchResult, error) {
	if z.dialErr != nil {
		return nil, z.dialErr
	}

	req := &proto.SearchRequest{
		Query: query.QToProto(q),
		Opts:  opts.ToProto(),
	}

	resp, err := z.client.Search(ctx, req)
	if err != nil {
		return nil, convertError(err)
	}

	var repoURLS map[string]string      // We don't use repoURLs in Khulnasoft
	var lineFragments map[string]string // We don't use lineFragments in Khulnasoft

	return zoekt.SearchResultFromProto(resp, repoURLS, lineFragments), nil
}

// List lists repositories. The query `q` can only contain
// query.Repo atoms.
func (z *zoektGRPCClient) List(ctx context.Context, q query.Q, opts *zoekt.ListOptions) (*zoekt.RepoList, error) {
	if z.dialErr != nil {
		return nil, z.dialErr
	}

	req := &proto.ListRequest{
		Query: query.QToProto(q),
		Opts:  opts.ToProto(),
	}

	resp, err := z.client.List(ctx, req)
	if err != nil {
		return nil, convertError(err)
	}

	return zoekt.RepoListFromProto(resp), nil
}

func (z *zoektGRPCClient) Close()         {}
func (z *zoektGRPCClient) String() string { return z.endpoint }

// convertError translates gRPC errors to well-known Go errors.
func convertError(err error) error {
	if err == nil || err == io.EOF {
		return nil
	}

	if status.Code(err) == codes.DeadlineExceeded {
		return context.DeadlineExceeded
	}

	if status.Code(err) == codes.Canceled {
		return context.Canceled
	}

	return err
}
