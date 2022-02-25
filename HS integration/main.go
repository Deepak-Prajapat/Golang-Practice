package main

import (
	"flag"
	"fmt"
	srvr "hsintegration/server"
)

func main() {
	server := srvr.NewServer()
	//server.Config.HubSpotURL = flag.String("HubSpot_URL", "https://api.hubapi.com/contacts/", "HubSpot API URL")
	server.Config.WebPort = flag.Int("PORT", 3000, "Web server port")

	// database configuration
	server.Config.DbHost = flag.String("DB_HOST", "localhost", "Postgres host")
	server.Config.DbName = flag.String("DB_NAME", "testing", "Postgres database name")
	server.Config.DbUser = flag.String("DB_USER", "postgres", "Postgres user name")
	server.Config.DbPass = flag.String("DB_PASS", "root", "Postgres password")
	server.Config.DbPort = flag.Int("DB_PORT", 5432, "Postgres port")
	server.Config.DbMaxConnections = flag.Int("DB_MAX_CONNECTIONS", 100, "Maximum Postgres connections")
	server.Config.DbVerboseLogs = flag.Bool("DB_LOG", false, "thorough logs of db activity")

	//Salesforce Configuration
	server.Config.SfGrantType = flag.String("grant_type", "password", "salesforce grand type")
	server.Config.SfSecurityToken = flag.String("SF_SECURITY_TOKEN", "ISv9Qophfz2JWOu67ojDofou", "salesforce security token")
	server.Config.SfTokenURL = flag.String("SF_Token", "https://login.salesforce.com/services/oauth2/token", "token url")
	server.Config.SfUsername = flag.String("SF_USERNAME", "golang@concret.io", "org username")
	server.Config.SfPassword = flag.String("SF_PASSWORD", "gopractice1", "salesforce org password")
	server.Config.SfClientID = flag.String("SF_CLIENT_ID", "3MVG9pRzvMkjMb6kDmaqfQNHc1l4BirqlwI4LccPebZ99s.xkiwceaoGm9ynZOIf2cZw3u3cFP3pdjJkVsoxR", "CONNECTED APP CLIENT ID")
	server.Config.SfClientSecret = flag.String("SF_CLIENT_SECRET", "A0E7E4128AC06CEFB0D7F34BC7CD4ACCC3125F688A638CA1A81FCF8F5B7FCB0F", "salesforce client secret")
	server.Config.SfInstanceURL = flag.String("INSTANCE_URL", "https://concretio89-dev-ed.my.salesforce.com", "ORG URL")

	flag.Parse()
	fmt.Println("----------------------------------------------------------:")

	//3MVG9pRzvMkjMb6kDmaqfQNHc1vDXo92FLCjQ.1p3zyVJmXSDeRB2y_p_Dw0ftZifNH9yVj.e7JS5qY3OXvqG
	//C7A0511142B7F523AC11690002B518F07009325292166C8CFCC1D4FFC28158EA

	fmt.Println("db password:", server.Config.SfPassword)
	fmt.Println("db password value:", *server.Config.SfPassword)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
