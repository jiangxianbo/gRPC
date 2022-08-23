package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ParamReq struct {
	X, Y int
}
type ParamRes struct {
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatal(err)
	}

	req := ParamReq{4, 2}

	var ret ParamRes

	// 调用方法
	err = conn.Call("Param.Multiply", req, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d * %d = %d\n", req.X, req.Y, ret.Pro)

	err = conn.Call("Param.Divide", req, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d / %d 商 %d，余数 = %d\n", req.X, req.Y, ret.Quo, ret.Rem)
}
