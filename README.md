# Experiments in UDP Multicasting

I have long been fascinated by broadcast networks. In the biological context, broadcast networks (and eavesdropping) are a central feature of anuran and insect choruses.

In TCP/IP there is a similar mechanism provided for in the UDP broadcast and multicast.

The following is rephrased from [here](https://www.quora.com/What-is-the-difference-between-broadcasting-and-multicasting).

**Broadcast** sends packets to all devices on a LAN. Unlike multicasting, there is no mechanism to limit the recipients of a broadcast: all packets go to all devices whether they want them or not. There is no mechanism to forward broadcasts between LANs.

**Multicast** sends packets to all devices in a specified group. Membership in a group is set up when devices send "join" packets to an upstream router, and routers and switches keep track of this membership. When multicast packets arrive at a switch, they are only sent to devices or segments (such as WiFi) where at least one device wants them. Multicast can traverse the networks where it has been configured.

## References

* https://www.quora.com/What-is-the-difference-between-broadcasting-and-multicasting
