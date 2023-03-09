package models

type QpInfoResponse struct {
	QpResponse
	Server *QpWhatsappServer `json:"server,omitempty"`
}

func (source *QpInfoResponse) ParseSuccess(server *QpWhatsappServer) {
	source.QpResponse.ParseSuccess("follow server information")
	source.Server = server
}
