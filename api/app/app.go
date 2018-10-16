package app

import (
	"github.com/astaxie/beego"
	"github.com/mdiazp/sirel-server/api/models"
	"github.com/mdiazp/sirel-server/api/models/models2"
	"github.com/mdiazp/sirel-server/api/pkg/authproviders"
	"github.com/mdiazp/sirel-server/api/pkg/authproviders/ldap"
	"github.com/mdiazp/sirel-server/api/pkg/authproviders/xxx"
	"github.com/mdiazp/sirel-server/api/pkg/cryptoutil"
)

var (
	model        models2.Model
	crypto       *cryptoutil.JWTCrypt
	ldapProvider authproviders.Provider
	xxxProvider  authproviders.Provider
)

func InitApp() {
	model = models.NewModel()
	crypto = cryptoutil.NewJWTCrypt()
	ldapProvider = ldap.GetProvider(
		beego.AppConfig.String("AdAddress"),
		beego.AppConfig.String("AdSuff"),
		beego.AppConfig.String("AdBDN"),
		beego.AppConfig.String("AdUser"),
		beego.AppConfig.String("AdPassword"),
	)
	xxxProvider = xxx.GetProvider()
}

func Model() models2.Model {
	return model
}

func Crypto() *cryptoutil.JWTCrypt {
	return crypto
}

const (
	AuthProviderLdap  = "ldap"
	AuthProviderSIREL = "xxx"
)

func AuthProvider(t string) authproviders.Provider {
	switch t {
	case AuthProviderLdap:
		return ldapProvider
	default:
		return xxxProvider
	}
}
