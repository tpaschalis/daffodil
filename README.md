# daffodil
Daffodil is a distributed ID generator, inspired by [Snowflake](https://github.com/twitter/snowflake) and [IdGen](https://github.com/RobThree/IdGen).

It can be either imported as a package, or be immediately deployed in your Kubernetes cluster.

## How it works

Daffodil generates 64-bit numbers, which can be used as roughly time-ordered UIDs, consisting of four parts :
- 1 bit unused sign bit
- 39 timestamp bits
- 16 node ID bits
- 8 sequence bits

Here's the breakdown of a daffodil-generated ID

```
Generated ID : 17801782472864612

0 000000000111111001111101010000001111000 1010100111110011 01100100
^              ^                          ^                ^
1-bit          39-bit                     16-bit           8-bit
sign           timestamp                  node ID          sequence

timestamp : 1061068920 ticks passed since epoch
node ID   : 34507
sequence  : 100th ID generated in this tick
```

The first *sign* bit is unused to simplify ordering in systems using both int/uint variables.   
The timestamp is defined in terms of *ticks* since a value configured as an *epoch*. The current tick length is 10 milliseconds.
If two IDs are generated within the same tick, the sequence is incremented.

This means that daffodil can generate up to 2^8=256 IDs per tick, per machine, and be deployed in up to 2^16=65536 machines.


## Benchmarks

## Getting started

### Installation
All you have to do is
```
go get github.com/tpaschalis/daffodil
```

### Usage

## Roadmap
- Provide a way to fine-tune the length of each ID component
- Provide a way to insert custom ID components

## License
`daffodil` is available under the MIT license. See the [LICENSE](https://github.com/tpaschalis/daffodil/blob/master/LICENSE) file for details.