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
2022/12/16 13:52:25 I also frob widgets and call AdditionalFunc()
2022/12/16 13:52:25 I'm not tied to a specific implementation but I need to be called when Canada frobs widgets
2022/12/16 13:52:25 Event -  I send emails about 🍁 Canada things
2022/12/16 13:52:25 Oh Canada! Oh Canada!
$ go build -tags us
$ ./build-flag-demo
2022/12/16 13:52:30 Event -  I send Emails about 🦅 MURICA things!!
2022/12/16 13:52:30 OH SAY CAN YOU SEEEEEEEEEE, BY THE DAWNS EARLY LIIIIGHT!
$ go build
$ ./build-flag-demo
panic: I can't be used on my own

goroutine 1 [running]:
main.init.0()
        /Users/nickglynn/Projects/build-flag-demo/events_default.go:7 +0x27
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
