package object

import (
	"fmt"
	"simple-casdoor/util"
)

type AdminProvider struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk unique" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Type         string `xorm:"varchar(100)" json:"type"`
	ClientId     string `xorm:"varchar(100)" json:"clientId"`
	ClientSecret string `xorm:"varchar(2000)" json:"clientSecret"`
	SignName     string `xorm:"varchar(100)" json:"signName"`
	TemplateCode string `xorm:"varchar(100)" json:"templateCode"`
}

func GetProvider(id string) (*AdminProvider, error) {
	owner, name := util.GetOwnerAndNameFromId(id)
	return getProvider(owner, name)
}

func AddProvider(provider *AdminProvider) (bool, error) {
	if provider.Owner == "" {
		provider.Owner = "fireboom"
	}
	affected, err := adapter.Engine.Insert(provider)
	if err != nil {
		return false, err
	}

	return affected != 0, nil
}

func UpdateProvider(provider *AdminProvider) (rows int64, err error) {
	if p, err := getProvider("fireboom", "provider_sms"); err != nil {
		return 0, err
	} else if p == nil {
		return 0, fmt.Errorf("provider does not exist")
	}

	affected, err := adapter.Engine.Table("provider").Where("owner=? and name=?", provider.Owner, provider.Name).Update(provider)
	if err != nil {
		return 0, err
	}

	return affected, nil
}

func getProvider(owner string, name string) (*AdminProvider, error) {
	if owner == "" || name == "" {
		return nil, nil
	}

	provider := AdminProvider{Owner: owner, Name: name}
	existed, err := adapter.Engine.Get(&provider)
	if err != nil {
		return &provider, err
	}

	if existed {
		return &provider, nil
	} else {
		return nil, nil
	}
}
