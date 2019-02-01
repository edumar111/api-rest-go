package main

type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}
type Movies []Movie

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}
