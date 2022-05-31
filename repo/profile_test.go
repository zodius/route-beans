package repo_test

import (
	"route-beans/repo"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestProfileRepo(t *testing.T) {
	suite.Run(t, new(ProfileRepoTestSuite))
}

type ProfileRepoTestSuite struct {
	suite.Suite
}

func (s *ProfileRepoTestSuite) TestParseFile() {
	repo := repo.NewProfileRepo()
	profile, err := repo.LoadProfileFile("../test-data/correct.yml")
	s.Assert().NoError(err)
	s.Assert().Equal("example", profile.Name)
}

func (s *ProfileRepoTestSuite) TestParseFileGatewayFormatError() {
	repo := repo.NewProfileRepo()
	_, err := repo.LoadProfileFile("../test-data/gateway_format_error.yml")
	s.Assert().Error(err)
}

func (s *ProfileRepoTestSuite) TestParseFileRuleFormatError() {
	repo := repo.NewProfileRepo()
	_, err := repo.LoadProfileFile("../test-data/rule_format_error.yml")
	s.Assert().Error(err)
}

func (s *ProfileRepoTestSuite) TestParseFileGatewayNotDefined() {
	repo := repo.NewProfileRepo()
	_, err := repo.LoadProfileFile("../test-data/gateway_not_define.yml")
	s.Assert().Error(err)
}
