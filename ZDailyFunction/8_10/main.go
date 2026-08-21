package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

func main() {
	/*
		context.WithTimeOut
		所属包：context
		功能：基于父Context创建一个带有超时时间的子Context，当到达指定时间或调用返回cancel函数时，该Context会自动取消
		适用场景：限制HTTP请求、数据库查询、RPC调用或长时间运行任务的超时时间，防止资源泄露或卡死
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("任务超时前完成")
	case <-ctx.Done():
		fmt.Println("任务超时或者取消", ctx.Err())
	}

	// ------------------------------------------------------------
	/*
		sync.OnceFunc
		所属包：sync
		功能：返回一个新的函数，保证传入的闭包函数仅被执行一次，相比传统的sync.Once，语法更加简洁直观
		适用场景：全局单例初始化、延迟适配加载】一次性资源连接
	*/
	initialize := sync.OnceFunc(
		func() { fmt.Println("只执行一次的初始化配置...") }) // 传一个函数给他，就会把这个函数锁定只执行一次

	// 多次调用，只有第一次生效
	initialize()
	initialize()
	initialize()

	// ---------------------------------------------------------------
	/*
		errors.Is
		所属包：errors
		功能：检查错误链中是否包含错误目标，能够递归解包包装后的错误（例如使用 fmt.Errorf("...: %w", err) 包装的错误）
		适用场景：在错误被多次包装后，准确判断底层是否为特定的 sentinel 错误（如 io.EOF、sql.ErrNoRows 等）
	*/
	err := fmt.Errorf("读取数据失败: %w", io.EOF) // 模拟包装错误
	if errors.Is(err, io.EOF) {             // 判断错误链中是否包含 io.EOF
		fmt.Println("检测到 EOF 错误，已到达文件末尾")
	} else {
		fmt.Println("其他未知错误:", err)
	}
}
