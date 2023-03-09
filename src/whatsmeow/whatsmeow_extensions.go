package whatsmeow

import (
	"encoding/base64"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"

	whatsapp "github.com/sufficit/sufficit-quepasa/whatsapp"
	whatsmeow "go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

type WhatsmeowLogLevel string

const (
	DebugLevel WhatsmeowLogLevel = "DEBUG"
	InfoLevel  WhatsmeowLogLevel = "INFO"
	WarnLevel  WhatsmeowLogLevel = "WARN"
	ErrorLevel WhatsmeowLogLevel = "ERROR"
)

func GetMediaTypeFromAttachment(source *whatsapp.WhatsappAttachment) whatsmeow.MediaType {
	return GetMediaTypeFromString(source.Mimetype)
}

// Traz o MediaType para download do whatsapp
func GetMediaTypeFromString(Mimetype string) whatsmeow.MediaType {

	msgType := whatsapp.GetMessageType(Mimetype)

	switch msgType {
	case whatsapp.ImageMessageType:
		return whatsmeow.MediaImage
	case whatsapp.AudioMessageType:
		return whatsmeow.MediaAudio
	case whatsapp.VideoMessageType:
		return whatsmeow.MediaVideo
	default:
		return whatsmeow.MediaDocument
	}
}

func ToWhatsmeowMessage(source whatsapp.IWhatsappMessage) (msg *waProto.Message, err error) {
	messageText := source.GetText()

	if !source.HasAttachment() {
		internal := &waProto.ExtendedTextMessage{Text: &messageText}
		msg = &waProto.Message{ExtendedTextMessage: internal}
	}

	return
}

func NewWhatsmeowMessageAttachment(response whatsmeow.UploadResponse, attach *whatsapp.WhatsappAttachment, media whatsmeow.MediaType) (msg *waProto.Message) {
	switch media {
	case whatsmeow.MediaImage:
		msg = &waProto.Message{ImageMessage: &waProto.ImageMessage{
			Url:           proto.String(response.URL),
			DirectPath:    proto.String(response.DirectPath),
			MediaKey:      response.MediaKey,
			FileEncSha256: response.FileEncSHA256,
			FileSha256:    response.FileSHA256,
			FileLength:    proto.Uint64(response.FileLength),

			Mimetype: proto.String(attach.Mimetype),
			Caption:  proto.String(attach.FileName),
		},
		}
		return
	case whatsmeow.MediaAudio:
		internal := &waProto.AudioMessage{
			Url:           proto.String(response.URL),
			DirectPath:    proto.String(response.DirectPath),
			MediaKey:      response.MediaKey,
			FileEncSha256: response.FileEncSHA256,
			FileSha256:    response.FileSHA256,
			FileLength:    proto.Uint64(response.FileLength),

			Mimetype: proto.String(attach.Mimetype),
			Ptt:      proto.Bool(ShouldUsePtt(attach.Mimetype)),
		}
		msg = &waProto.Message{AudioMessage: internal}
		return
	case whatsmeow.MediaVideo:
		internal := &waProto.VideoMessage{
			Url:           proto.String(response.URL),
			DirectPath:    proto.String(response.DirectPath),
			MediaKey:      response.MediaKey,
			FileEncSha256: response.FileEncSHA256,
			FileSha256:    response.FileSHA256,
			FileLength:    proto.Uint64(response.FileLength),

			Mimetype: proto.String(attach.Mimetype),
			Caption:  proto.String(attach.FileName),
		}
		msg = &waProto.Message{VideoMessage: internal}
		return
	default:
		internal := &waProto.DocumentMessage{
			Url:           proto.String(response.URL),
			DirectPath:    proto.String(response.DirectPath),
			MediaKey:      response.MediaKey,
			FileEncSha256: response.FileEncSHA256,
			FileSha256:    response.FileSHA256,
			FileLength:    proto.Uint64(response.FileLength),

			Mimetype: proto.String(attach.Mimetype),
			FileName: proto.String(attach.FileName),
		}
		msg = &waProto.Message{DocumentMessage: internal}
		return
	}
}

// Use that to set if the message should be sent as PTT audio
func ShouldUsePtt(Mimetype string) bool {
	return strings.Contains(Mimetype, "ogg") && strings.Contains(Mimetype, "opus")
}

func GetStringFromBytes(bytes []byte) string {
	if bytes != nil {
		return base64.StdEncoding.EncodeToString(bytes)
	}
	return ""
}
