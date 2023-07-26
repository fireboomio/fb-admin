package global

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	Wg.Add(1)
	go func() {
		Mq <- body.Response.Status
		Wg.Done()
		Wg.Wait()
	}()

	return body.Response, nil
}
