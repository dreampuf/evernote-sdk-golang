# Evernote SDK Golang

This project was simple code generated from [Evernote-thrift](https://github.com/evernote/evernote-thrift)-1.25 .

# Simple

TODO

# How to generate yourself code

1. Install the newest Thrift. It's a type alias bug of golang generator [THRIFT-2955](https://issues.apache.org/jira/browse/THRIFT-2955).
1. Clone the official evernote-thrift repo `https://github.com/evernote/evernote-thrift`
1. Generator with this command:

    thrift -strict -nowarn --allow-64bit-consts --allow-neg-keys --gen go:package_prefix=github.com/dreampuf/evernote-sdk-golang/ evernote-thrift/src/UserStore.thrift

1. Enjoy!
