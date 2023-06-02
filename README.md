# entity-naming

Entity-naming is a Go library for generating names of OpenConfig entities
backed by a set of vendor-specific implementations. These entities include
physical and virtual interfaces and components like linecards and fabrics.

Entity-naming is self-contained set of side-effect-free utility functions.
The library does _not_ communicate with  devices or external network entities.
It contains implementations of the vendor-specific naming conventions alone,
presented through a uniform ventor-neutral API, so that client code can
retrieve the name of the entity without implementing vendor-specific logic.

## Contributions

Contributions are more than welcome, specially from the vendors themselves.

To add support for a new vendor, your PR should introduce a new directory
named for that vendor under
[internal](https://github.com/openconfig/entity-naming/tree/main/internal).
