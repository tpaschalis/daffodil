# daffodil
Daffodil is a distributed ID generator, inspired by [Snowflake](https://github.com/twitter/snowflake) and [IdGen](https://github.com/RobThree/IdGen).

It can be either imported as a package, or be immediately deployed in your Kubernetes cluster.

## How it works

Daffodil generates 64-bit numbers, which can be used as roughly time-ordered UIDs, consisting of four parts :
- 1 bit unused sign bit
- 39 timestamp bits
- 16 node ID bits
- 8 sequence bits

Here's the breakdown of a generated ID

```
Generated ID : 17801782472864612

0 000000000111111001111101010000001111000 1010100111110011 01100100
^              ^                          ^                ^
1-bit          39-bit                     16-bit           8-bit
sign           timestamp                  node ID          sequence

timestamp : 1061068920 ticks passed since epoch
node ID   : 43507
sequence  : 100th ID generated in this tick
```

The first *sign* bit is unused to simplify ordering in systems using both int/uint variables.   
The timestamp is defined in terms of *ticks* since a value configured as an *epoch*. The current tick length is 10 milliseconds.
If two IDs are generated within the same tick, the sequence is incremented.

This means that daffodil can generate up to 2^8=256 IDs per tick, per machine, and be deployed in up to 2^16=65536 machines.

Daffodil depends on the system clock to calculate the number of ticks since the epoch. There are some checks in place to ensure the monotonic nature of the clock, but using something like NTP to keep all distributed nodes synchronized will help maintain the order-ability of the generated IDs and lessen the chance of collisions.

## Benchmarks

To get some magnitude perspective of a modern system at scale, as of May 2020, Twitter generates ~10k tweets per second. As we mentioned, daffodil should be able to generate 256 IDs per 10msec *tick* i.e. handle ~25600 rps.

My first attempt was to simply `go run` a single instance on a beefy laptop and measure performance with `httperf` and `ab`. The application was able to sustain bursts of up to 20 k requests per second, until it started replying with non-2xx responses. When trying to pick up the pace, I either run [out of file descriptors](https://stackoverflow.com/questions/29484074/httperf-file-descriptor-limit) or [ephemeral ports](https://stackoverflow.com/questions/1216267/ab-program-freezes-after-lots-of-requests-why), so I figured, let's move to a more realistic testing scenario, in a Kubernetes deployment!

```
$ ab -n 30000 -c 30 "http://localhost:8080/"
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
...
Time taken for tests:   1.475 seconds
Complete requests:      30000
Failed requests:        3
...
Requests per second:    20336.63 [#/sec] (mean)
Time per request:       1.475 [ms] (mean)
Time per request:       0.049 [ms] (mean, across all concurrent requests)
...
```

**TODO: Add K8S benchmarks**

## Getting started

First off, we need to initialize a *Config* struct and *Daffodil* itself.
```go
cfg, err := daffodil.NewConfig()
if err != nil {
	log.Fatal(err)
}

df, err = daffodil.NewDaffodil(cfg)
if err != nil {
	log.Fatal(err)
}
```

Currently, the only configured value is the Node ID, which by default is calculated from the  host's private IP, to avoid conflicts between nodes.   

Setting `${DAFFODIL_NODEID_MODE}` to `HOSTNAME` means the hostname (eg. pod name) will be translated into a 16-but uint NodeID.   

Similarly, setting `${DAFFODIL_NODEID_MODE}` to `CUSTOM` will read the environment variable stored in `${DAFFODIL_NODEID_CUSTOM}` and use it as the Node ID.

After initializing `daffodil`, you can just use the `Next()` method to produce IDs.
```go
type ID uint64

func (d *Daffodil) Next() (ID, error)
```

In cases that the application encounters an error with the clock or the sequence overflows, the method will return an error.

### Installation
All you have to do is
```
go get github.com/tpaschalis/daffodil
```

### Usage

`daffodil` includes a simple standalone implementation, which you can `go build` to deploy a binary, a [Docker image](https://github.com/tpaschalis/daffodil/blob/master/Dockerfile), and a [Kubernetes deployment](https://github.com/tpaschalis/daffodil/blob/master/deployment.yaml).

The standalone application runs on port `:8080` and contains two routes;   
`/` the root url which simply generates an ID and    
`/dismantle?id=` which is used to disassemble an ID to its components
```
$ cd cmd/daffodil 

$ go run main.go &
[1] 18938 

$ curl "localhost:8080/"
17908816211997440

$ curl "localhost:8080/dismantle?id=17908816211997440"
{
    "id": 17908816211997440,
    "node": 43507,
    "sequence": 0,
    "timestamp": 1067448628
}
```

The docker image can be run using
```
$ docker build --tag daffodil .
$ docker run --publish 8080:8080 --detach --name daffodil daffodil
39b02d55beef7254e08b1f0492f24e3c0461d415ddd8bd94237834492349845f
```

The Kubernetes deployment ...

## Roadmap
- Provide a way to fine-tune the length of each ID component
- Provide a way to insert custom ID components

## License
`daffodil` is available under the MIT license. See the [LICENSE](https://github.com/tpaschalis/daffodil/blob/master/LICENSE) file for details.
