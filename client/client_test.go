package client

import (
	"context"
	"github.com/dreampuf/evernote-sdk-golang/edam"
	"os"
	"testing"
	"time"
)

var (
	yes = true
	EvernoteKey = os.Getenv("KEY")
	EvernoteSecret = os.Getenv("SECRET")
	EvernoteAuthorToken = os.Getenv("TOKEN")
)


func TestClient(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(15) * time.Second)
	c := NewClient(EvernoteKey, EvernoteSecret, SANDBOX)
	us, err := c.GetUserStore()
	if err != nil {
		t.Fatal(err)
	}
	userUrls, err := us.GetUserUrls(ctx, EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	ns, err := c.GetNoteStoreWithURL(userUrls.GetNoteStoreUrl())
	if err != nil {
		t.Fatal(err)
	}
	notebook, err := ns.GetDefaultNotebook(ctx, EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	if notebook == nil {
		t.Fatal("Invalid Note")
	}
	// optional tag filter
	filterTags := []edam.GUID{}
	tags, err := ns.ListTags(ctx, EvernoteAuthorToken)
	if err != nil {
		t.Fatal(err)
	}
	for _, tag := range tags {
		filterTags = append(filterTags, tag.GetGUID())
	}
	noteMetadataList, err := ns.FindNotesMetadata(ctx, EvernoteAuthorToken, &edam.NoteFilter{
		//Ascending:                    &yes,
		//TagGuids:                     filterTags,
	}, 0, 1000, &edam.NotesMetadataResultSpec{
		IncludeTitle:               &yes,
		IncludeContentLength:       &yes,
		IncludeCreated:             &yes,
		IncludeUpdated:             &yes,
		IncludeTagGuids:            &yes,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("total note: %d\n", noteMetadataList.GetTotalNotes())
	for n, noteMate := range noteMetadataList.GetNotes() {
		t.Logf("%d - %s - %s\n", n, time.Unix(int64(noteMate.GetCreated())/1000, 0), noteMate.GetTitle())
	}
}

func TestRequestToken(t *testing.T) {
	/*
	PRODUCTION: evernote production
	SANDBOX: evernote sandbox
	YINXIANG: yinxiangbiji
	 */
	c := NewClient(EvernoteKey, EvernoteSecret, SANDBOX)
	callBackURL := "http://YOUR_SERVER_CALL_BACK_URL"
	requestToken, url, err := c.GetRequestToken(callBackURL)
	if err != nil {
		t.Fatal(err)
	}
	if requestToken == nil {
		t.Fatal("Failed token request")
	}
	if len(url) < 1 {
		t.Fatal("Invalid URL")
	}

	// in the call back handler
	// if you are using gin-gonic https://github.com/gin-gonic/gin
	oauthToken := "OAUTH_TOKEN" // c.Query("oauth_token")
	oauthVerifier := "OAUTH_VERIFIER" // c.Query("oauth_verifier")
	_ = oauthToken

	accessToken, err := c.GetAuthorizedToken(requestToken, oauthVerifier)
	if err == nil {
		us, _ := c.GetUserStore()
		userUrls, _ := us.GetUserUrls(context.Background(), accessToken.Token)
		// ... referring the test case of TestClient
		_ = userUrls
	}
}