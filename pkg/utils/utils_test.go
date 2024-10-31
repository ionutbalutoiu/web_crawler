package utils_test

import (
	"net"
	"testing"
	"web_crawler/pkg/utils"

	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
}

func (s *UtilsTestSuite) TestGetHostDomain() {
	testCases := []struct {
		Name     string
		Host     string
		Expected string
		Err      error
	}{
		{
			Name:     "Simple_Host",
			Host:     "example.com",
			Expected: "example.com",
			Err:      nil,
		},
		{
			Name:     "Host_With_Port",
			Host:     "example.com:8080",
			Expected: "example.com",
			Err:      nil,
		},
		{
			Name:     "Invalid_Host",
			Host:     "example.com:8080:9090",
			Expected: "",
			Err:      &net.AddrError{Err: "too many colons in address", Addr: "example.com:8080:9090"},
		},
	}
	for _, tc := range testCases {
		s.Run(tc.Name, func() {
			domain, err := utils.GetHostDomain(tc.Host)
			s.Equal(tc.Err, err)
			s.Equal(tc.Expected, domain)
		})
	}
}

func TestUtilsTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UtilsTestSuite))
}
