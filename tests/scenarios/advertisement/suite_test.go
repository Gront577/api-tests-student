package advertisement

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/managers/auth"
	authModels "api-tests-template/internal/managers/auth/models"

	base "api-tests-template/tests"
)

// TestSuite вЂ” РЅР°Р±РѕСЂ С‚РµСЃС‚РѕРІ РґР»СЏ РїСЂРѕРІРµСЂРєРё СЃРѕР·РґР°РЅРёСЏ РѕР±СЉСЏРІР»РµРЅРёР№.
type TestSuite struct {
	suite.Suite
	loginData authModels.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("РђРІС‚РѕСЂРёР·Р°С†РёСЏ РїРѕР»СЊР·РѕРІР°С‚РµР»СЏ СЃ РєСЂРµРґР°РјРё РёР· РїРµСЂРµРјРµРЅРЅС‹С… РѕРєСЂСѓР¶РµРЅРёСЏ")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) SetupTest() {
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}
