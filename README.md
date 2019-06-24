# ipSearch
Send ping to all link local addresses to identify the IP of the terminal.

### Requirements
- golang 1.11.4
- [go-ping](https://github.com/sparrc/go-ping)

### Option
- **-c uint** Upper limit of parallel processing (default 100)
- **-debug** Display all results
- **-t uint** Set of ping timeout [ms] (default 100)

### Note
- If the timeout period you set is too short (< 15 ms), you may not be able to find it.
