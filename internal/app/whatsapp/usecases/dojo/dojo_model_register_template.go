package dojo

import (
	"encoding/json"
)

const jsonModel = `{
   "name":"seasonal_promotion_text_only",
   "language":"en | pt_BR",
   "category":"MARKETING | UTILITY",
   "components": {
      "header": {
         "type":"HEADER",
         "format":"TEXT",
         "text":"Our [1] is on!",
         "example":{
            "header_text":[
               "Summer Sale"
            ]
         }
      },
      "body": {
         "type":"BODY",
         "text":"Shop now through [1] and use code [2] to get [3] off of all merchandise.",
         "example":{
            "body_text":[
               [
                  "the end of August",
                  "25OFF",
                  "25%"
               ]
            ]
         }
      },
      "footer": {
         "type":"FOOTER",
         "text":"Use the buttons below to manage your marketing subscriptions"
      },
      "buttons": {
         "type":"BUTTONS",
         "buttons":[
            {
               "type":"QUICK_REPLY",
               "text":"Unsubcribe from Promos"
            },
            {
               "type":"QUICK_REPLY",
               "text":"Unsubscribe from All"
            }
         ]
      }
   }
}`

type ModelRegisterTemplate struct{}

func NewModelRegisterTemplate() *ModelRegisterTemplate {
	return &ModelRegisterTemplate{}
}

func (temp *ModelRegisterTemplate) ModelRegisterTemplate() any {
	var model interface{}
	_ = json.Unmarshal([]byte(jsonModel), &model)

	return model
}
