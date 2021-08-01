# sparkplugclidecoder

---
Sparkplug cli decoder allow you to debug in command line a stream of messages received in the standard input in
hexadecimal format. Currently only tested with mosquitto_sub

## Usage

```shell
usage: sparkplugclidecoder [<flags>]

Flags:
--help                Show context-sensitive help (also try --help-long and --help-man).
-s, --sep="--"            to separate the messages currently: msg<newline>sep<newline>
-f, --format="prototext"  Output format: prototext, protojson
```

## Example
Tells mosquitto_sub to write the sparkplug message as hexadecimal string so sparkplugclidecoder
reads from it and prints the decoded in prototext

```shell
mosquitto_sub -L mqtt://localhost:1883/testtopic/# -F %x | ./sparkplugclidecoder
```


