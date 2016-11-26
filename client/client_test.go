package client
import (
	"testing"
)
const (
	EvernoteKey string = ""
	EvernoteSecret string = ""
	EvernoteAuthorToken string = ""
)


func TestClient(t *testing.T) {
	c := NewClient(EvernoteKey, EvernoteSecret, SANDBOX)
	us, err := c.GetUserStore()
	if err != nil {
		t.Fatal(err)
	}
	url, err := us.GetNoteStoreUrl(EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	if len(url) < 1 {
		t.Fatal("Invalid URL")
	}
	ns, err := c.GetNoteStoreWithURL(url)
	if err != nil {
		t.Fatal(err)
	}
	note, err := ns.GetDefaultNotebook(EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	if note == nil {
		t.Fatal("Invalid Note")
	}
}

func TestRequestToken(t *testing.T) {
	c := NewClient(EvernoteKey, EvernoteSecret, SANDBOX)
	requestToken, url, err := c.GetRequestToken("http://test/")
	if err != nil {
		t.Fatal(err)
	}
	if requestToken == nil {
		t.Fatal("Failed token request")
	}
	if len(url) < 1 {
		t.Fatal("Invalid URL")
	}
}
