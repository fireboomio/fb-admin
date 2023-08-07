package DeleteOne

import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func PostResolve(hook *base.HookRequest, body generated.System__Role__DeleteOneBody) (res generated.System__Role__DeleteOneBody, err error) {
	hook.Logger().Info("PostResolve")

	return body, nil
}
