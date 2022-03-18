package models

type AllBackendServerResponse struct {
	Servers []Info `json:"servers"`
}

type Info struct {
	Url     string `json:"name"`
	Healthy bool   `json:"healthy"`
}
