# snowflake-cli

Easily format Discord and Twitter Snowflakes.

### Usage

```bash
# Build project, or run `go build -o snowflake .`
make

# Format default
./snowflake 1191470808191737927
# 1704139177130 - 2024-01-01T20:59:37+01:00

# Format in a custom timezone
TZ=America/New_York ./snowflake -s 1191470808191737927 -f "%u - %p (East Coast)"
# 1704139177130 - 2024-01-01T14:59:37-05:00 (East Coast)

# Format custom output
./snowflake -s 1191470808191737927 -f "%u (%t:R) meow %%uwu%%"
# 1704139177130 (<t:1704139177:R>) meow %uwu%
```

### Options

```bash
Usage of ./snowflake:
  -f string
    	Format string, options: %u (unix ms), %p (RFC3339), %t (<t:unix>), %t:S (<t:unix:style>) (default "%u - %p")
  -s int
    	Snowflake to parse
  -twitter
    	Set to parse Twitter snowflake instead of Discord
```

#### Custom Format

Please see the [Discord Timestamp Style docs](https://discord.com/developers/docs/reference#message-formatting-timestamp-styles) for more info on the `%t` and `%t:S` format.