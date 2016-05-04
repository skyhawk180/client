// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/mobile.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type HellokbfsArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Echo      string `codec:"echo" json:"echo"`
}

type MobileInterface interface {
	Hellokbfs(context.Context, HellokbfsArg) (string, error)
}

func MobileProtocol(i MobileInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.mobile",
		Methods: map[string]rpc.ServeHandlerDescription{
			"hellokbfs": {
				MakeArg: func() interface{} {
					ret := make([]HellokbfsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]HellokbfsArg)
					if !ok {
						err = rpc.NewTypeError((*[]HellokbfsArg)(nil), args)
						return
					}
					ret, err = i.Hellokbfs(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type MobileClient struct {
	Cli rpc.GenericClient
}

func (c MobileClient) Hellokbfs(ctx context.Context, __arg HellokbfsArg) (res string, err error) {
	err = c.Cli.Call(ctx, "keybase.1.mobile.hellokbfs", []interface{}{__arg}, &res)
	return
}