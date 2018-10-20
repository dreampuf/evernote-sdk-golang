# Evernote SDK Golang

This project was simple code generated from [Evernote-thrift](https://github.com/evernote/evernote-thrift)-1.29.

# Simple

```bash
go get -u github.com/dreampuf/evernote-sdk-golang/...
```

See [client_test.go](client/client_test.go)

```golang
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
```

# How to generate yourself code

1. Install the newest Thrift.
1. Clone the official evernote-thrift repo `https://github.com/evernote/evernote-thrift`
1. Generator with this command:

     ```bash
     thrift -strict -nowarn --allow-64bit-consts --allow-neg-keys --gen go:package_prefix=github.com/dreampuf/evernote-sdk-golang/,thrift_import=github.com/apache/thrift/lib/thrift -strict -nowarn --allow-64bit-consts --allow-neg-keys --gen go:package_prefix=github.com/dreampuf/evernote-sdk-golang/,thrift_import=github.com/apache/thrift/lib/go/thrift -I src/ -r --out github.com/dreampuf/evernote-sdk-golang src/UserStore.thrift
     ```

1. Enjoy!

