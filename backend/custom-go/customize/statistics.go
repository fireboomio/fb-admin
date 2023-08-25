package customize

import (
	"custom-go/generated"
	"custom-go/pkg/plugins"
	"github.com/graphql-go/graphql"
)

type (
	queryRawIn  = generated.Statistics__QueryRawInternalInput
	queryRawOut = generated.Statistics__QueryRawResponseData
)

const (
	queryRawPath = generated.Statistics__QueryRaw
)

func init() {
	plugins.RegisterGraphql(&Statistics_schema)
}

var Statistics_schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"GetVisitTrending": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "GetVisitTrending",
					Fields: graphql.Fields{
						"days":  &graphql.Field{Type: graphql.String},
						"count": &graphql.Field{Type: graphql.Int},
					},
				})),
				Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
					grc, _, err := plugins.ResolveArgs[any](p)
					if err != nil {
						return
					}

					jsonVal, err := plugins.ExecuteInternalRequestMutations[queryRawIn, queryRawOut](grc.InternalClient, queryRawPath,
						queryRawIn{Query: "SELECT date(visitedAt) days, COUNT(id) count from case_visitlog where visitedAt BETWEEN '2020-01-01' AND '2020-12-31' GROUP BY days;"})
					if err != nil {
						return
					}
					res = jsonVal.Main_queryRaw
					return
				},
			},
			"GetMonthlySales": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "GetMonthlySales",
					Fields: graphql.Fields{
						"months":     &graphql.Field{Type: graphql.String},
						"totalSales": &graphql.Field{Type: graphql.Float},
					},
				})),
				Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
					grc, _, err := plugins.ResolveArgs[any](p)
					if err != nil {
						return
					}

					jsonVal, err := plugins.ExecuteInternalRequestMutations[queryRawIn, queryRawOut](grc.InternalClient, queryRawPath,
						queryRawIn{Query: "SELECT date(day) months, SUM(sales) totalSales from case_salelog where day BETWEEN '2019-10-01' AND '2020-09-30' GROUP BY months;"})
					if err != nil {
						return
					}
					res = jsonVal.Main_queryRaw
					return
				},
			},
			"GetSalesTop10": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "GetSalesTop10",
					Fields: graphql.Fields{
						"shopName":   &graphql.Field{Type: graphql.String},
						"totalSales": &graphql.Field{Type: graphql.Float},
					},
				})),
				Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
					grc, _, err := plugins.ResolveArgs[any](p)
					if err != nil {
						return
					}

					jsonVal, err := plugins.ExecuteInternalRequestMutations[queryRawIn, queryRawOut](grc.InternalClient, queryRawPath,
						queryRawIn{Query: "SELECT shopName, SUM(sales) totalSales from case_salelog GROUP BY shopName ORDER BY totalSales DESC;"})
					if err != nil {
						return
					}

					res = jsonVal.Main_queryRaw
					return
				},
			},
			"GetSaleTypePercent": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "GetSaleTypePercent",
					Fields: graphql.Fields{
						"typeId":     &graphql.Field{Type: graphql.Int},
						"typeName":   &graphql.Field{Type: graphql.String},
						"totalSales": &graphql.Field{Type: graphql.Float},
					},
				})),
				Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
					grc, _, err := plugins.ResolveArgs[any](p)
					if err != nil {
						return
					}

					jsonVal, err := plugins.ExecuteInternalRequestMutations[queryRawIn, queryRawOut](grc.InternalClient, queryRawPath,
						queryRawIn{Query: "SELECT case_saletype.id typeId, case_saletype.name typeName,SUM(sales) totalSales from case_salelog, case_saletype WHERE case_salelog.typeId = case_saletype.id GROUP BY case_salelog.typeId ORDER BY totalSales DESC;"})
					if err != nil {
						return
					}

					res = jsonVal.Main_queryRaw
					return
				},
			},
		},
	}),
})
