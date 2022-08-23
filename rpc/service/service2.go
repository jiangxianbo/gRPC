package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Param struct {
}

type ParamRequest struct {
	X, Y int
}

type ParamResponse struct {
	Pro int
	Quo int
	Rem int
}

func (p Param) Multiply(req ParamRequest, res *ParamResponse) error {
	res.Pro = req.X * req.Y
	return nil
}

func (p Param) Divide(req ParamRequest, res *ParamResponse) error {
	res.Quo = req.X / req.Y
	res.Rem = req.X % req.Y
	return nil
}

func main() {
	// 1.注册服务
	rect := new(Param)
	rpc.Register(rect)
	// 2.服务绑定
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe("127.0.0.1:8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
