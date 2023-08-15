package global

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"strconv"
)

func OnOriginResponse(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
	//result, _ := plugins.ExecuteInternalRequestQueries[getIsOpenI, getIsOpenO](hook.InternalClient, generated.System__Log__GetIsOpen, getIsOpenI{})

	//if result.Main_findFirstdic.IsOpen {
	go func() {
		defer Wg.Done()
		if strconv.Itoa(body.Response.StatusCode) != "" {
			Mq <- strconv.Itoa(body.Response.StatusCode)
		}

		hook.Logger().Info("qwertyuiopcvbnm")
	}()
	Wg.Wait()
	//}

	return body.Response, nil
}
