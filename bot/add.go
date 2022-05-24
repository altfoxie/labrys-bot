package bot

import (
	"fmt"
	"html"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/altfoxie/labrys-bot/i18n"
	"github.com/altfoxie/labrys-bot/storage"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

// Handler for /addvoice command.
func (b *Bot) onAddVoice(locale *i18n.Locale, message *telego.Message, arg string) error {
	if arg == "" {
		return lo.T2(b.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				locale.NoName,
			),
		)).B
	}

	var fileID string
	if reply := message.ReplyToMessage; reply != nil {
		switch {
		case reply.Voice != nil:
			fileID = reply.Voice.FileID
		case reply.Audio != nil:
			fileID = reply.Audio.FileID
		case reply.Video != nil:
			fileID = reply.Video.FileID
		case reply.VideoNote != nil:
			fileID = reply.VideoNote.FileID
		case reply.Document != nil:
			fileID = reply.Document.FileID
		}
	}
	if fileID == "" {
		return lo.T2(b.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				locale.NoReply,
			),
		)).B
	}

	tgFile, err := b.GetFile(&telego.GetFileParams{FileID: fileID})
	if err != nil {
		return err
	}

	temp := path.Join(os.TempDir(), fmt.Sprintf("labrys-voice-%d.ogg", time.Now().UnixNano()))
	defer os.Remove(temp)

	cmd := exec.Command("ffmpeg",
		"-hide_banner",
		"-v", "error",
		"-i", fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.Token(), tgFile.FilePath),
		"-map_metadata", "-1",
		"-metadata", "artist=Labrys",
		"-metadata", "album=Labrys",
		"-metadata", "author=Labrys",
		"-metadata", "title=@LabrysBot",
		"-c:a", "libopus",
		temp,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		logrus.WithError(err).Error("Error occurred while converting voice.")
		return err
	}

	file, err := os.Open(temp)
	if err != nil {
		logrus.WithError(err).Error("Error occurred while opening converted voice.")
		return err
	}
	defer file.Close()

	sent, err := b.SendVoice(
		tu.Voice(
			tu.ID(lo.Ternary(b.channelID != 0, b.channelID, message.Chat.ID)),
			tu.File(file),
		).WithCaption(arg),
	)
	if err != nil {
		return err
	}

	if err := b.storage.Voices.Create(storage.Voice{
		Name:      arg,
		FileID:    sent.Voice.FileID,
		MessageID: lo.Ternary(b.channelID != 0, sent.MessageID, 0),
	}); err != nil {
		return err
	}

	return lo.T2(b.SendMessage(
		tu.Message(
			tu.ID(message.Chat.ID),
			locale.VoiceAdded,
		),
	)).B

}

// Handler for /addpasta command.
func (b *Bot) onAddPasta(locale *i18n.Locale, message *telego.Message, arg string) error {
	if arg == "" {
		return lo.T2(b.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				locale.NoName,
			),
		)).B
	}

	text := ""
	if reply := message.ReplyToMessage; reply != nil {
		text, _ = lo.Coalesce(reply.Text, reply.Caption)
	}
	if text == "" {
		return lo.T2(b.SendMessage(
			tu.Message(
				tu.ID(message.Chat.ID),
				locale.NoReply,
			),
		)).B
	}

	var sentID int
	if b.channelID != 0 {
		if sent, err := b.SendMessage(tu.Message(
			tu.ID(b.channelID),
			fmt.Sprintf("<b>%s</b>\n\n%s", arg, html.EscapeString(text)),
		).WithParseMode(telego.ModeHTML)); err == nil {
			sentID = sent.MessageID
		}
	}

	if err := b.storage.Pastas.Create(storage.Pasta{
		Name:      arg,
		Content:   text,
		MessageID: sentID,
	}); err != nil {
		return err
	}

	return lo.T2(b.SendMessage(
		tu.Message(
			tu.ID(message.Chat.ID),
			locale.PastaAdded,
		),
	)).B
}
