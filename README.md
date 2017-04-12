# Experiments in UDP Multicasting

I have long been fascinated by broadcast networks. In the biological context, broadcast networks (and eavesdropping) are a central feature of anuran and insect choruses.

In TCP/IP there is a similar mechanism provided for in the UDP broadcast and multicast. This package wraps Golang's UDP functions in the `net` package.

## Background

The following is rephrased from [here](https://www.quora.com/What-is-the-difference-between-broadcasting-and-multicasting).

**Broadcast** sends packets to all devices on a LAN. Unlike multicasting, there is no mechanism to limit the recipients of a broadcast: all packets go to all devices whether they want them or not. There is no mechanism to forward broadcasts between LANs.

**Multicast** sends packets to all devices in a specified group. Membership in a group is set up when devices send "join" packets to an upstream router, and routers and switches keep track of this membership. When multicast packets arrive at a switch, they are only sent to devices or segments (such as WiFi) where at least one device wants them. Multicast can traverse the networks where it has been configured.

## Examples

This package comes with some small command line utilities in the `examples` dir.

In a terminal window, run the following from the root of this repository.

```bash
$ go run examples/pinger/main.go
# => Broadcasting to 239.0.0.0:9999
```

In a separate terminal window, run the following, also from the root of this repository.

```bash
$ go run examples/listener/main.go
# => Listening on 239.0.0.0:9999
# 2017/04/12 12:53:24 13 bytes read from 192.168.1.129:51335
# 2017/04/12 12:53:24 00000000  68 65 6c 6c 6f 2c 20 77  6f 72 6c 640a           |hello, world.|
#
# 2017/04/12 12:53:25 13 bytes read from 192.168.1.129:51335
# 2017/04/12 12:53:25 00000000  68 65 6c 6c 6f 2c 20 77  6f 72 6c 64 0a           |hello, world.|
```

You may run as many instances of `listeners` or `pingers` as you would like. Concurrency is handled by the router (or switch).

## References

The code in this repository is derived from [a Gist](https://gist.github.com/fiorix/9664255) created by [Andre Fiori](https://gist.github.com/fiorix)

* https://www.quora.com/What-is-the-difference-between-broadcasting-and-multicasting
* https://groups.google.com/forum/#!msg/golang-nuts/nbmYWwHCgPc/ZBw2uH6Bdi4J
* https://en.wikipedia.org/wiki/Multicast_address
* http://www.tldp.org/HOWTO/Multicast-HOWTO-2.html
* https://support.mcommstv.com/hc/en-us/articles/202306226-Choosing-multicast-addresses-and-ports
