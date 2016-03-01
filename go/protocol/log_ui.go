// Auto-generated by avdl-compiler v1.0.3 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/log_ui.avdl
//   Generated : Sun Feb 28 2016 22:03:20 GMT-0500 (EST)

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type LogArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Level     LogLevel `codec:"level" json:"level"`
	Text      Text     `codec:"text" json:"text"`
}

type LogUiInterface interface {
	Log(context.Context, LogArg) error
}

func LogUiProtocol(i LogUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.logUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"log": {
				MakeArg: func() interface{} {
					ret := make([]LogArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]LogArg)
					if !ok {
						err = rpc.NewTypeError((*[]LogArg)(nil), args)
						return
					}
					err = i.Log(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LogUiClient struct {
	Cli rpc.GenericClient
}

func (c LogUiClient) Log(ctx context.Context, __arg LogArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.logUi.log", []interface{}{__arg}, nil)
	return
}