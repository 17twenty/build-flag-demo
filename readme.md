# Demo of Using Build Flags in Go

There's different approaches to doing region specific deployments.
One avenue that is wells supported is to use Go Build tags.

This is where we use Go interfaces to define our expected behaviour
and then the region specific is pushed into build specific files.
You can use negative flags in your builds to ensure that we have a panicing
version to indicate poor builds.

```bash

$ go build -tags ca
$ ./build-flag-demo 
2022/12/16 12:46:31 I also frob widgets and call AdditionalFunc()
2022/12/16 12:46:31 Event -  I send emails about 🍁 Canada things
$ go build -tags us
$ ./build-flag-demo
2022/12/16 12:46:56 Event -  I send Emails about 🦅 MURICA things!!
$ go build
$ ./build-flag-demo 
2022/12/16 10:51:33 Not Implemented
panic: Not Implemented

goroutine 1 [running]:
log.Panic({0xc000092f40?, 0xc000092f70?, 0x1004f19?})
        /usr/local/Cellar/go/1.19.2/libexec/src/log/log.go:388 +0x65
main.EventEngine.raiseEvent(...)
        /build-flag-demo/events_default.go:11
main.main()
        /build-flag-demo/main.go:7 +0x48
```

## Alternative Options

An alternative way to do this at runtime rather than build time could be:

```go
...
country := os.Getenv("COUNTRY_CODE")

var CountrySpecific foo = null

switch country {
    case "US":
        // create US specific thing
        ...
        foo = America{}
    case "AU":
        // create AU specific thing
        ...
        foo = Aussie{}
}

CountrySpecific.DoSomething()

}

```

The downside here is you'd need to do this check for all scenarios
Due to being runtime only, poor hygiene can cause bugs where checks
have been removed/moved or the behavior changed which is hard to catch.
