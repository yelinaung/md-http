md-http
=======

Serve your markdown file in browser.

It doesn't do anything but render your given markdown file in the browser.

Installation
-------------
- install go either via [gvm](https://github.com/moovweb/gvm) or [download](http://golang.org/doc/install) from website.
- `go get github.com/yelinaung/md-http`
- `cd $GOPATH/src/github.com/yelinaung/md-http && go build main.go`
- You can add it to $PATH by `go install`
- Go to your folder where the markdown files are. 
- Run `md-http`
- Open Browser and go to `localhost:3000/[your-markdown-file-name].md`

Contributing
------------

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

License
-------
MIT
