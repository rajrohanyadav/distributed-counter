# Distributed Counter

```bash
> make run
go build
./counters
==============================
G-Counter - Grow only counter
==============================
Before Merge
client 1: 3
client 2: 6
After Merge
client 1: 9
client 2: 9
==============================
PN-Counter - Positive Negative Counter
==============================
Before Merge
client 3: 1
client 4: 3
After Merge
client 3: 4
client 4: 4
```
