package tcp_server_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var stringFormat = regexp.MustCompile(`\S+`)
var sessFormat = regexp.MustCompile(`dnn=([^\,]+),sd=([^\,]+),sst=(\S+)`)

func TestRegexp(t *testing.T) {
	{
		tmp := stringFormat.FindAllString("reg", -1)
		fmt.Printf("%q", tmp)
	}
	{
		tmp := stringFormat.FindAllString("reg 127.0.0.1", -1)
		fmt.Printf("%q", tmp)
	}
	{
		tmp := stringFormat.FindAllString("sess 2 add", -1)
		fmt.Printf("%q", tmp)
	}
	{
		tmp := stringFormat.FindAllString("sess 2 del", -1)
		fmt.Printf("%q", tmp)
	}
	{
		tmp := stringFormat.FindAllString("dereg", -1)
		fmt.Printf("%q", tmp)
	}

	{
		tmp := stringFormat.FindAllString("show all", -1)
		fmt.Printf("%q", tmp)
	}
	{
		tmp := stringFormat.FindAllString("show 4", -1)
		fmt.Printf("%q", tmp)
	}
	{
		tmp := sessFormat.FindStringSubmatch("dnn=internet,sst=1,sd=010203")
		if assert.NotNil(t, tmp) {
			fmt.Printf("%q", tmp)
		}
	}

}
