package mandrill
/*
import(
	"time"
	"net/mail"
)

type Time struct {
	Time time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.UTC().Format("2006-01-02 15:04:05")), nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.Parse("2006-01-02 15:04:05", string(data))
	return err
}

type Address struct {
	Address *mail.Address
}

func (a Address) MarshalJSON() ([]byte, error) {
	return []byte(a.Address.Address), nil
}

func (a *Address) UnMarshalJSON(data []byte) (err error) {
	a.Address, err = mail.ParseAddress(string(data))
	return err
}
*/


type Variable struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
