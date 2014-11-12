package inq

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor/conf/inq/ast"
	"github.com/mperham/inspeqtor/conf/inq/lexer"
	"github.com/mperham/inspeqtor/conf/inq/parser"
	"github.com/stretchr/testify/assert"
)

func TestMysqlParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/mysql.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.ProcessCheck)
	assert.Equal(t, check.Name, "mysql")
	assert.Equal(t, len(check.Parameters), 0)
	assert.Equal(t, len(check.Rules), 4)
	assert.Equal(t, check.Rules[1].Threshold.Parsed, 50)
	assert.Equal(t, check.Rules[1].Threshold.Raw, "50")
	assert.Equal(t, check.Rules[1].Threshold.PerSec, false)

	assert.Equal(t, check.Rules[2].Threshold.Parsed, 1024)
	assert.Equal(t, check.Rules[2].Threshold.Raw, "1k/sec")
	assert.Equal(t, check.Rules[2].Threshold.PerSec, true)

	assert.Equal(t, check.Rules[3].Threshold.Parsed, 2)
	assert.Equal(t, check.Rules[3].Threshold.Raw, "2/sec")
	assert.Equal(t, check.Rules[3].Threshold.PerSec, true)
}

func TestBadAmount(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/bad_amount.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	assert.Nil(t, obj)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "Invalid amount"))
}

func TestBasicServiceParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/memcached.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.ProcessCheck)
	assert.Equal(t, check.Name, "memcached")
	assert.Equal(t, len(check.Rules), 2)
	assert.Equal(t, len(check.Parameters), 4)
	assert.Equal(t, check.Parameters["owner"], "dev")
	assert.Equal(t, check.Parameters["foo"], "bar")
	assert.Equal(t, check.Parameters["endpoint"], "/foo")
	assert.Equal(t, check.Parameters["quoted"], "whoa sp\"aces")
	assert.Equal(t, check.Rules[0].Actions[1].Name(), "alert")
	assert.Equal(t, check.Rules[0].Actions[2].Name(), "reload")
	assert.Equal(t, check.Rules[1].Metric.Family, "cpu")
	assert.Equal(t, check.Rules[1].Metric.Name, "user")
}

func TestBasicHostParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/host.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.HostCheck)
	assert.Equal(t, len(check.Rules), 3)
}
