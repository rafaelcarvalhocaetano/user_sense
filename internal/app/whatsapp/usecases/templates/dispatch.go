package templates

import (
	"botwhatsapp/internal/app/whatsapp/usecases/templates/dto"
	"botwhatsapp/internal/interfaces/services"
	"fmt"
	"strconv"
	"strings"
)

type ExecuteDispatch struct {
	dispatch services.WTAGateway
}

func NewExecuteDispatch(d services.WTAGateway) *ExecuteDispatch {
	return &ExecuteDispatch{dispatch: d}
}

func (dd *ExecuteDispatch) ExecuteDispatch(input *dto.Dispatch) ([]map[string]any, error) {
	var response []map[string]any
	for _, c := range input.Clients {
		cleanValueStr := strings.ReplaceAll(c.Value, ".", "")
		amount, _ := strconv.Atoi(cleanValueStr)
		message := dto.Message{
			MessagingProduct: "whatsapp",
			RecipientType:    "individual",
			To:               c.Phone,
			Type:             "template",
			Template: dto.Template{
				Name:     input.TemplateName,
				Language: dto.Language{Code: "pt_BR"},
				Components: []dto.Component{
					{
						Type: "body",
						Parameters: []dto.Parameter{
							{
								Type: "text",
								Text: c.Name,
							},
							{
								Type: "currency",
								Currency: &dto.Currency{
									FallbackValue: fmt.Sprintf("%v", c.Value),
									Code:          "BRL",
									Amount1000:    amount,
								},
							},
							{
								Type: "date_time",
								DateTime: &dto.DateTime{
									FallbackValue: c.Datetime,
									DayOfWeek:     5,
									Year:          2024,
									Month:         6,
									DayOfMonth:    6,
									Hour:          0,
									Minute:        0,
									Calendar:      "GREGORIAN",
								},
							},
						},
					},
				},
			},
		}

		resp, err := dd.dispatch.Send("messages", message)
		if err != nil {
			response = append(response, map[string]any{"usuário": c.Name, "error_message": err.Error()})
		}
		response = append(response, map[string]any{"usuário": c.Name, "success": resp})
	}

	return response, nil
}
