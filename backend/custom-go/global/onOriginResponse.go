package global

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	return body.Response, nil
}
