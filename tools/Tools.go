package tools

import (
	"fmt"
	"github.com/scarletcc/weworkapi_go/api"
	"strings"
	"regexp"
	"io/ioutil"
	"os"
	"time"
)

func GenerateCorpAPIMethods() time.Time {
	now := time.Now()
	re := regexp.MustCompile("(^\\w)|(_\\w)")
	content := fmt.Sprintf(`package api
/*
 * Generated CorpAPI methods
 * %v
 */
`, now)
	tmpl := `
func (c *CorpAPI) %v(args map[string]interface{}) (map[string]interface{}, error) {
	return c.HttpCall(CORP_API_TYPE["%v"], args)
}
		`
	for key, _ := range api.CORP_API_TYPE {
		if key == "GET_ACCESS_TOKEN" {
			continue
		}
		funcName := re.ReplaceAllStringFunc(strings.ToLower(key), func(str string) string {
			return strings.ToUpper(strings.Replace(str, "_", "", -1))
		})
		content = content + fmt.Sprintf(tmpl, funcName, key)
	}

	ioutil.WriteFile("../api/CorpAPIMethod.go", []byte(content), os.ModePerm)
	return now
}