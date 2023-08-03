package global

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"strconv"
)

func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	//等待协程处理完成
	Wg.Add(1)
	//获取处理完的code
	go func() {
		if strconv.Itoa(body.Response.StatusCode) != "" {
			Mq <- strconv.Itoa(body.Response.StatusCode)
		}
		Wg.Done()
		Wg.Wait()
	}()

	return body.Response, nil
}
