package client

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/dreampuf/evernote-sdk-golang/edam"
	"github.com/mrjones/oauth"
)

type EnvironmentType int

const (
	SANDBOX EnvironmentType = iota
	PRODUCTION
	YINXIANG
)

type EvernoteClient struct {
	host        string
	oauthClient *oauth.Consumer
	userStore   *edam.UserStoreClient
}

func NewClient(key, secret string, envType EnvironmentType) *EvernoteClient {
	host := "www.evernote.com"
	if envType == SANDBOX {
		host = "sandbox.evernote.com"
	} else if envType == YINXIANG {
		host = "app.yinxiang.com"
	}
	client := oauth.NewConsumer(
		key, secret,
		oauth.ServiceProvider{
			RequestTokenUrl:   fmt.Sprintf("https://%s/oauth", host),
			AuthorizeTokenUrl: fmt.Sprintf("https://%s/OAuth.action", host),
			AccessTokenUrl:    fmt.Sprintf("https://%s/oauth", host),
		},
	)
	return &EvernoteClient{
		host:        host,
		oauthClient: client,
	}
}

func (c *EvernoteClient) GetRequestToken(callBackURL string) (*oauth.RequestToken, string, error) {
	return c.oauthClient.GetRequestTokenAndUrl(callBackURL)
}

func (c *EvernoteClient) GetAuthorizedToken(requestToken *oauth.RequestToken, oauthVerifier string) (*oauth.AccessToken, error) {
	return c.oauthClient.AuthorizeToken(requestToken, oauthVerifier)
}

func (c *EvernoteClient) GetUserStore() (*edam.UserStoreClient, error) {
	if c.userStore != nil {
		return c.userStore, nil
	}
	evernoteUserStoreServerURL := fmt.Sprintf("https://%s/edam/user", c.host)
	thriftTransport, err := thrift.NewTHttpClient(evernoteUserStoreServerURL)
	thriftClient := thrift.NewTStandardClient(thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(thriftTransport), thrift.NewTBinaryProtocolFactory(true, true).GetProtocol(thriftTransport))
	if err != nil {
		return nil, err
	}
	c.userStore = edam.NewUserStoreClient(thriftClient)
	return c.userStore, nil
}

func (c *EvernoteClient) GetNoteStore(ctx context.Context, authenticationToken string) (*edam.NoteStoreClient, error) {
	us, err := c.GetUserStore()
	if err != nil {
		return nil, err
	}
	userUrls, err := us.GetUserUrls(ctx, authenticationToken)
	if err != nil {
		return nil, err
	}

	ns, err := c.GetNoteStoreWithURL(userUrls.GetNoteStoreUrl())
	return ns, nil
}

func (c *EvernoteClient) GetNoteStoreWithURL(notestoreURL string) (*edam.NoteStoreClient, error) {
	thriftTransport, err := thrift.NewTHttpClient(notestoreURL)
	thriftClient := thrift.NewTStandardClient(thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(thriftTransport), thrift.NewTBinaryProtocolFactory(true, true).GetProtocol(thriftTransport))
	if err != nil {
		return nil, err
	}
	return edam.NewNoteStoreClient(thriftClient), nil
}
