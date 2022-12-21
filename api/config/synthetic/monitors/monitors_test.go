package monitors_test

import (
	"encoding/json"
	"testing"

	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors"
)

var input = `{
    "script": {
        "events": [
            {
                "description": "jhjhjh",
                "type": "navigate",
                "url": "https://www.orf.at",
                "validate": [
                    {
                        "failIfFound": false,
                        "isRegex": true,
                        "match": "kkl",
                        "target": {
                            "window": "k"
                        },
                        "type": "text_match"
                    }
                ],
                "wait": {
                    "timeoutInMilliseconds": 60000,
                    "validation": {
                        "failIfFound": false,
                        "match": "kjkj",
                        "target": {
                            "locators": [
                                {
                                    "type": "css",
                                    "value": "jjj"
                                }
                            ]
                        }
                    }
                }
            }
        ],
        "type": "clickpath",
        "version": "1.0"
    },
    "tags": [],
    "type": "BROWSER"
}`

func TestJSON(t *testing.T) {
	data := []byte(input)
	monitor := new(monitors.BrowserSyntheticMonitorUpdate)
	if err := json.Unmarshal(data, &monitor); err != nil {
		t.Error(err)
	}
}
