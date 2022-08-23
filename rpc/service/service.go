package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {
}

func (r Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

func (r Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height * p.Width) * 2
	return nil
}

func main() {
	// 1.注册服务
	rect := new(Rect)
	// 注册一个rect的服务
	rpc.Register(rect)
	// 2.服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3.监听服务
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Panicln(err)
	}
}
