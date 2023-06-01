package response

type ResponseMeta struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}
type Userparam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Success struct {
	ResponseMeta
	Data Userparam `json:"data"`
}
type SuccessGetAll struct {
	ResponseMeta
	Data []Userparam `json:"data"`
}
