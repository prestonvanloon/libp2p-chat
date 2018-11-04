# Chat 

This is an experiment on connecting clients via relay nodes with a common DHT.

# Goals 

- Clients connect via relay [DONE]
- Clients discover each other in a common DHT via their relay address [DONE]

Example workflow:

1) Client A connects to a DHT of some kind.
1) Client A provides their connection address with (and without) relay path.
1) Client A requests peers from DHT, DHT returns empty list (or list of only cl
Client A).
1) Client B connects to DHT.
1) Client B provides their connection addresses (w/ & w/o relay).
1) Client B requests peers from DHT, DHT returns list with Client A.
1) Client B attempts to connect to Client A, prioritizing non-relay addresses.
1) Client A is unreachable from Client B without the relay node.
1) Client B is connected to Client A via relay node.
