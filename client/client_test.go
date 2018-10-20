package client
import (
	"context"
	"testing"
	"time"
)
const (
	EvernoteKey string = ""
	EvernoteSecret string = ""
	EvernoteAuthorToken string = ""
)


func TestClient(t *testing.T) {
	clientCtx, _ := context.WithTimeout(context.Background(), time.Duration(15) * time.Second)
	c := NewClient(EvernoteKey, EvernoteSecret, SANDBOX)
	us, err := c.GetUserStore()
	if err != nil {
		t.Fatal(err)
	}
	userUrls, err := us.GetUserUrls(clientCtx, EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	ns, err := c.GetNoteStoreWithURL(userUrls.GetNoteStoreUrl())
	if err != nil {
		t.Fatal(err)
	}
	note, err := ns.GetDefaultNotebook(clientCtx, EvernoteAuthorToken)
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