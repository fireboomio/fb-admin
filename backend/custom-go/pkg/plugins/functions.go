package plugins

import (
	"custom-go/pkg/base"
	"custom-go/pkg/consts"
	"custom-go/pkg/utils"
	"custom-go/pkg/wgpb"
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/invopop/jsonschema"
	"github.com/labstack/echo/v4"
	"github.com/tidwall/sjson"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	schemaRefPrefix       = "#/$defs/"
	swaggerRefPrefix      = "#/definitions/"
	DefinitionRefProperty = "$defs"
)

func RegisterFunction[I, O any](hookFunc func(*base.HookRequest, *base.OperationBody[I, O]) (*base.OperationBody[I, O], error), conf ...*HookConfig) {
	callerName := utils.GetCallerName(consts.FUNCTIONS)
	apiPrefixPath := "/" + consts.FUNCTIONS
	apiPath := path.Join(apiPrefixPath, callerName)

	base.AddEchoRouterFunc(func(e *echo.Echo) {
		e.Logger.Debugf(`Registered hookFunction [%s]`, apiPath)
		e.POST(apiPath, buildOperationHook(callerName, consts.FUNCTIONS, ConvertBodyFunc[I, O](hookFunc), func(in, out *base.OperationBody[any, any]) {
			in.Response = out.Response
		}))
	})

	base.AddHealthFunc(func(e *echo.Echo, s string, report *base.HealthReport) {
		operation := &wgpb.Operation{
			Name:          callerName,
			Path:          apiPath,
			OperationType: wgpb.OperationType_MUTATION,
		}
		if len(conf) > 0 {
			operation.AuthenticationConfig = &wgpb.OperationAuthenticationConfig{AuthRequired: conf[0].AuthRequired}
			operation.LiveQueryConfig = &wgpb.OperationLiveQueryConfig{Enabled: conf[0].EnableLiveQuery}
			operation.AuthorizationConfig = conf[0].AuthorizationConfig
			operation.OperationType = conf[0].OperationType
		}

		var i I
		var o O
		// 解析入参和出参
		inputSchema := jsonschema.Reflect(i)
		outputSchema := jsonschema.Reflect(o)

		operation.VariablesSchema = BuildSchema(inputSchema)
		operation.ResponseSchema = BuildSchema(outputSchema)

		operationBytes, err := json.Marshal(operation)
		if err != nil {
			e.Logger.Errorf("json marshal failed, err: %v", err.Error())
			return
		}

		err = os.WriteFile(filepath.Join(consts.FUNCTIONS, callerName)+consts.JSON_EXT, operationBytes, 0644)
		if err != nil {
			e.Logger.Errorf("write file failed, err: %v", err.Error())
			return
		}

		report.Functions = append(report.Functions, callerName)
	})
}

func BuildSchema(schema *jsonschema.Schema) string {
	defs := make(openapi3.Schemas)
	for name, internalSchema := range schema.Definitions {
		defs[name] = parseJsonschemaToSwaggerSchema(internalSchema)
	}

	refStr := strings.TrimPrefix(schema.Ref, "#/$defs/")
	schemaRef := defs[refStr]
	delete(defs, refStr)

	bytes, _ := json.Marshal(schemaRef)

	res := string(bytes)

	if len(defs) == 0 {
		return res
	}
	res, _ = sjson.Set(res, DefinitionRefProperty, defs)
	return res
}

func parseJsonschemaToSwaggerSchema(schema *jsonschema.Schema) (result *openapi3.SchemaRef) {
	if schema.Ref != "" {
		result = openapi3.NewSchemaRef(strings.ReplaceAll(schema.Ref, schemaRefPrefix, swaggerRefPrefix), nil)
		return
	}

	result = &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:     schema.Type,
		Format:   schema.Format,
		Default:  schema.Default,
		Title:    schema.Title,
		Required: schema.Required,
	}}

	if schema.Items != nil {
		result.Value.Items = parseJsonschemaToSwaggerSchema(schema.Items)
		return
	}

	for _, item := range schema.OneOf {
		result.Value.OneOf = append(result.Value.OneOf, parseJsonschemaToSwaggerSchema(item))
	}
	for _, item := range schema.AnyOf {
		result.Value.AnyOf = append(result.Value.AnyOf, parseJsonschemaToSwaggerSchema(item))
	}
	for _, item := range schema.AllOf {
		result.Value.AllOf = append(result.Value.AllOf, parseJsonschemaToSwaggerSchema(item))
	}

	if properties := schema.Properties; properties != nil {
		result.Value.Properties = make(openapi3.Schemas)
		for _, key := range properties.Keys() {
			itemSchema, _ := properties.Get(key)
			result.Value.Properties[key] = parseJsonschemaToSwaggerSchema(itemSchema.(*jsonschema.Schema))
		}
	}
	return
}
