package logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type FormatterSuite struct {
	suite.Suite
}

func (suite *FormatterSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
}

func TestFormatterSuite(t *testing.T) {
	suite.Run(t, new(FormatterSuite))
}

func (suite *FormatterSuite) TestFormatter() {
	oldOut := log.StandardLogger().Out
	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	hostname, _ := os.Hostname()
	formatter := customFormatter{
		formatter: &log.TextFormatter{},
		additionalFields: map[string]string{
			"name":     "api",
			"hostname": hostname,
			"foo":      "bar",
		},
	}
	log.SetFormatter(formatter)

	log.Info("hello world")

	regex := `time=".*" level=info msg="hello world" foo=bar hostname=` + hostname + ` name=api`

	suite.Regexp(regex, buf.String())
	// restore log target
	log.SetOutput(oldOut)
}
