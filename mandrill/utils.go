package mandrill

import(
	"time"
	"net/mail"
)


const(
	DateFormat string = "2006-01-02 15:04:05"
)

func FormatTime(t time.Time) string {
	return t.UTC().Format(DateFormat)
}

func FormatEmail(m mail.Address) string {
	return m.Address
}
