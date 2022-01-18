# vtysock

vtysock is a vtysh replacement that directly sends commands to the vty sockets of the daemons.
By skipping the parsing and validation checks done in vtysh, vtysock can achieve a significant speed improvement when executing commands.

## usage

```
# vtysock -d ospfd -c 'show ip ospf neigh'
```
