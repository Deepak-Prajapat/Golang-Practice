package test

import (
	"github.com/unrolled/render"
	"hsintegration/salesforce"
	"hsintegration/server"
	u "hsintegration/utility"
)

func PrepareRendrer() {
	var r = render.New()
	u.R = r
}

func PrepareHsConfig() {
	s := server.NewServer()
	//s.Config.HubSpotURL = flag.String("HubSpot_URL", "https://api.hubapi.com/contacts/", "HubSpot API URL")
	//s.Config.WebPort = flag.Int("Test_PORT", 3000, "Web server port")
	var n = 3000
	s.Config.WebPort = &n
}

func PrepareSFConfig() {
	s := server.NewServer()
	//Salesforce Configuration
	//s.Config.SfGrantType = flag.String("Test_grant_type", "password", "salesforce grand type")
	//s.Config.SfSecurityToken = flag.String("Test_SF_SECURITY_TOKEN", "ISv9Qophfz2JWOu67ojDofou", "salesforce security token")
	//s.Config.SfTokenURL = flag.String("Test_SF_Token", "https://login.salesforce.com/services/oauth2/token", "token url")
	//s.Config.SfUsername = flag.String("Test_SF_USERNAME", "golang@concret.io", "org username")
	//s.Config.SfPassword = flag.String("Test_SF_PASSWORD", "gopractice1", "salesforce org password")
	//s.Config.SfClientID = flag.String("Test_SF_CLIENT_ID", "3MVG9pRzvMkjMb6kDmaqfQNHc1l4BirqlwI4LccPebZ99s.xkiwceaoGm9ynZOIf2cZw3u3cFP3pdjJkVsoxR", "CONNECTED APP CLIENT ID")
	//s.Config.SfClientSecret = flag.String("Test_SF_CLIENT_SECRET", "A0E7E4128AC06CEFB0D7F34BC7CD4ACCC3125F688A638CA1A81FCF8F5B7FCB0F", "salesforce client secret")
	//s.Config.SfInstanceURL = flag.String("Test_INSTANCE_URL", "https://concretio89-dev-ed.my.salesforce.com", "ORG URL")
	s.Config.SfGrantType = GetStringPtr("password")
	s.Config.SfSecurityToken = GetStringPtr("ISv9Qophfz2JWOu67ojDofou")
	s.Config.SfTokenURL = GetStringPtr("https://login.salesforce.com/services/oauth2/token")
	s.Config.SfUsername = GetStringPtr("golang@concret.io")
	s.Config.SfPassword = GetStringPtr("gopractice1")
	s.Config.SfClientID = GetStringPtr("3MVG9pRzvMkjMb6kDmaqfQNHc1l4BirqlwI4LccPebZ99s.xkiwceaoGm9ynZOIf2cZw3u3cFP3pdjJkVsoxR")
	s.Config.SfClientSecret = GetStringPtr("A0E7E4128AC06CEFB0D7F34BC7CD4ACCC3125F688A638CA1A81FCF8F5B7FCB0F")
	s.Config.SfInstanceURL = GetStringPtr("https://concretio89-dev-ed.my.salesforce.com")

	salesforce.SetupSFConfig(s.Config.SfPassword, s.Config.SfGrantType, s.Config.SfClientID,
		s.Config.SfClientSecret, s.Config.SfSecurityToken, s.Config.SfUsername)
}

func GetStringPtr(str string) *string {
	x := str
	return &x
}
