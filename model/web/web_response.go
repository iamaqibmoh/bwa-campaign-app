package web

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type WebResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
