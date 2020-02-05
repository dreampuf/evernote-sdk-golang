# Evernote SDK Golang

Unofficial SDK for Evernote. According official [Evernote-thrift](https://github.com/evernote/evernote-thrift) 1.29

# Installation

```bash
go get -u github.com/dreampuf/evernote-sdk-golang/...
```

# Example

1. Apply your [evernote developer key](https://dev.evernote.com/doc/) and [developer token](https://dev.evernote.com/doc/articles/dev_tokens.php) (an oauth access token for simply testing). More information: https://dev.evernote.com/doc/
1. Clone the library
    ```bash
   $ git clone https://github.com/dreampuf/evernote-sdk-golang.git
    ```
1. Run the test cases with your credentials
```bash
$ cd evernote-sdk-golang
$ KEY=YOUR_KEY SECRET=YOUR_SECRET TOKEN="YOUR_DEVERLOPER_TOKEN" go test ./...
ok  	github.com/dreampuf/evernote-sdk-golang/client	1.556s
?   	github.com/dreampuf/evernote-sdk-golang/edam	[no test files]
?   	github.com/dreampuf/evernote-sdk-golang/edam/note_store-remote	[no test files]
?   	github.com/dreampuf/evernote-sdk-golang/edam/user_store-remote	[no test files]
```


## As a library

Referring to [client_test.go](client/client_test.go)

```golang
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
```

# How to generate code step by step

1. Install the latest Thrift. *(`brew install thrift` if you are using MacOS)*
1. Clone evernote-thrift repo `https://github.com/evernote/evernote-thrift`
1. Generate golang version specs through this command:

     ```bash
     thrift -strict -nowarn \
       --allow-64bit-consts \
       --allow-neg-keys \
       --gen go:package_prefix=github.com/dreampuf/evernote-sdk-golang/,thrift_import=github.com/apache/thrift/lib/thrift \
       -strict -nowarn --allow-64bit-consts \
       --allow-neg-keys \
       --gen go:package_prefix=github.com/dreampuf/evernote-sdk-golang/,thrift_import=github.com/apache/thrift/lib/go/thrift \
       -I src/ -r \
       --out github.com/dreampuf/evernote-sdk-golang src/UserStore.thrift
     ```

1. There are some minor type convert issues you need to do a manually fix. I found Evernote internally has a Golang version SDK and here is what [their suggestion](https://github.com/evernote/evernote-thrift/issues/10#issuecomment-324966201). 

# License

MIT
