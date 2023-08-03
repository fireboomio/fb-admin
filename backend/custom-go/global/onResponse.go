package global

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"strconv"
)

func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	//
	Wg.Add(1)
	go func() {
		if strconv.Itoa(body.Response.StatusCode) != "" {
			Mq <- strconv.Itoa(body.Response.StatusCode)
		}
		Wg.Done()
		Wg.Wait()
	}()

	return body.Response, nil
}
