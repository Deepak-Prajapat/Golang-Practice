package main

import (
	"flag"
	"hsintegration/server"
)

func main() {
	s := server.NewServer()
	s.Config.WebPort = flag.Int("PORT", 3000, "Web s port")

	// database configuration
	s.Config.DbHost = flag.String("DB_HOST", "localhost", "Postgres host")
	s.Config.DbName = flag.String("DB_NAME", "testing", "Postgres database name")
	s.Config.DbUser = flag.String("DB_USER", "postgres", "Postgres user name")
	s.Config.DbPass = flag.String("DB_PASS", "root", "Postgres password")
	s.Config.DbPort = flag.Int("DB_PORT", 5432, "Postgres port")
	s.Config.DbMaxConnections = flag.Int("DB_MAX_CONNECTIONS", 100, "Maximum Postgres connections")
	s.Config.DbVerboseLogs = flag.Bool("DB_LOG", false, "thorough logs of db activity")

	//Salesforce Configuration
	s.Config.SfGrantType = flag.String("grant_type", "password", "salesforce grand type")
	s.Config.SfSecurityToken = flag.String("SF_SECURITY_TOKEN", "ISv9Qophfz2JWOu67ojDofou", "salesforce security token")
	s.Config.SfTokenURL = flag.String("SF_Token", "https://login.salesforce.com/services/oauth2/token", "token url")
	s.Config.SfUsername = flag.String("SF_USERNAME", "golang@concret.io", "org username")
	s.Config.SfPassword = flag.String("SF_PASSWORD", "gopractice1", "salesforce org password")
	s.Config.SfClientID = flag.String("SF_CLIENT_ID", "3MVG9pRzvMkjMb6kDmaqfQNHc1l4BirqlwI4LccPebZ99s.xkiwceaoGm9ynZOIf2cZw3u3cFP3pdjJkVsoxR", "CONNECTED APP CLIENT ID")
	s.Config.SfClientSecret = flag.String("SF_CLIENT_SECRET", "A0E7E4128AC06CEFB0D7F34BC7CD4ACCC3125F688A638CA1A81FCF8F5B7FCB0F", "salesforce client secret")
	s.Config.SfInstanceURL = flag.String("INSTANCE_URL", "https://concretio89-dev-ed.my.salesforce.com", "ORG URL")

	flag.Parse()
	if err := s.Run(); err != nil {
		panic(err)
	}
}
