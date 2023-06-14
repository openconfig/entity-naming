# entity-naming

Entity-naming is a Go library for computing the vendor-specific names of
OpenConfig entities. The goal of the library is to enable client code to
retrieve the name of network entities without including vendor-specific logic
itself. These entities include physical and virtual interfaces, components like
linecards and fabrics, and more.

The computation of the names is delegated to a set of vendor-specific naming
implementations. The functions provided by the library are self-contained and
free of any I/O. The code does not _not_ communicate with network devices or
external services. Rather, the library provides simple, pure-Go implementations
of the vendor-specific entity naming conventions.

## Example Usage

Here is the signature for the function for computing the name of an aggregate
interface:

```go
func AggregateInterface(dev *DeviceParams, index int) (string, error)
```

All of the functions provided by the library accept a set of of device
parameters, provided in a `DeviceParams` struct. To compute the name of a
Juniper PTX10008, for example, you would construct the following device
parameters:

```go
dev := &entname.DeviceParams{
    Vendor: naming.JUNIPER,
    HardwareModel: "PTX10008",
}
```

All index parameters accepted by the library are _zero-based_indices_, even in
cases where the vendor starts their numbering at 1 or later. For example, to
compute the name of the first aggregate interface, use the call

```go
aggName, err := AggregateInterface(dev, 0)
```

For the Juniper device parameters we provided, `aggName` will be "ae0", but for
an Arista device it will be "Port-Channel1", for Cisco "Bundle-Ether1", and for
Nokia "lag1."

## Traffic Queues

The library includes a `CommonTrafficQueues` function that returns a set of
common traffic service class tags to queues names. These traffic class tags are
defined as follows, from highest to lowest priority:

| Tag | QoS Level                     | Description                                                                                                                                        |
| --- | ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| NC1 | Network Control               | Traffic critical to the functionality of network devices and protocols, including routing protocol communication (OSPF, BGP, ISIS)                 |
| AF4 | Assured Forwarding Priority 4 | Critical application and production traffic that requires low latency communication                                                                |
| AF3 | Assured Forwarding Priority 3 | Mid-priority internal applications and higher bandwidth user applications that have lower latency sensitivity than AF4                             |
| AF2 | Assured Forwarding Priority 2 | Latency-insensitive, minimal loss-tolerating internal traffic that requires some amount of guaranteed delivery                                     |
| AF1 | Assured Forwarding Priority 1 | High-bandwidth, latency-insensitive traffic                                                                                                        |
| BE1 | Best Effort                   | Latency-insensitive, loss-insensitive traffic that can exhibit a substantial amount of packet loss and therefore should not carry any user traffic |
| BE0 | High-loss Best Effort         | Latency-insensitive, loss-insensitive traffic that can exhibit a higher loss rate than BE1                                                         |

## Contributions

Contributions are more than welcome, specially from the vendors themselves.

To add support for a new vendor, your PR should add a new value to the `Vendor`
enum in
[entname.go](https://github.com/openconfig/entity-naming/blob/main/entname/entname.go)
and add a new directory named for that vendor under
[internal](https://github.com/openconfig/entity-naming/tree/main/internal).
