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

## Examples
Tells mosquitto_sub to write the sparkplug message as hexadecimal string so sparkplugclidecoder
reads from it and prints the decoded in prototext

```shell
> mosquitto_sub -L mqtt://localhost:1883/testtopic/# -F %x | sparkplugclidecoder

timestamp:  1621956521303
metrics:  {
  name:  "Device Control/Scan Rate ms"
  timestamp:  1621956521303
  datatype:  3
  int_value:  6000
}
seq:  165
```

## Installation
```shell
go install github.com/dubyte/sparkplugclidecoder@latest
```

## Binary releases
- https://github.com/dubyte/sparkplugclidecoder/releases
