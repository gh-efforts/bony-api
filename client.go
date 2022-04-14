package bony_api

import (
	"context"
	"fmt"
	"github.com/filecoin-project/go-jsonrpc"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"net/http"
	"strings"
)

func NewBonyRPC(ctx context.Context, addr string, requestHeader http.Header) (API, jsonrpc.ClientCloser, error) {
	var res APIStruct
	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin",
		[]interface{}{
			&res.Internal,
			&res.StateAPIStruct.Internal,
			&res.ChainAPIStruct.Internal,
			&res.LotusServiceAPIStruct.Internal,
		},
		requestHeader,
	)
	return &res, closer, err
}

func GetAPI(ctx context.Context, addrStr string, token string) (API, jsonrpc.ClientCloser, error) {
	addrStr = strings.TrimSpace(addrStr)

	ainfo := cliutil.APIInfo{Addr: addrStr, Token: []byte(token)}

	addr, err := ainfo.DialArgs("v0")
	if err != nil {
		return nil, nil, fmt.Errorf("could not get DialArgs: %w", err)
	}

	return NewBonyRPC(ctx, addr, ainfo.AuthHeader())
}
