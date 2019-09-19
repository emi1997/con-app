package client

import (
	"fmt"
	"context"
)

//AddMapping lets you add mapping
func AddMapping() {
	var mappingTemplate = `
	{
		"dynamic_templates": [
		  {
			"integers": {
			  "match_mapping_type": "long",
			  "mapping": {
				"type": "integer"
			  }
			}
		  },
		  {
			"strings": {
			  "match_mapping_type": "string",
			  "mapping": {
				"type": "text",
				"fields": {
				  "raw": {
					"type":  "keyword",
					"ignore_above": 256
				  }
				}
			  }
			}
		  }
		]
	}
	`

	NewMapping, err := Client.PutMapping().Index("school").BodyString(mappingTemplate).Do(context.TODO())
	if err != nil {
		panic(err)
		// fmt.Printf("expected put mapping to succeed; got: %v", err)
	}

	//fmt.Println(mappingTemplate, NewMapping)
	if !NewMapping.Acknowledged {
		fmt.Print("something went wrong")
	}
}
