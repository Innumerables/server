package response

import "server/model/example"

type CustomerResponse struct {
	Customer example.Customer `json:"customer"`
}
