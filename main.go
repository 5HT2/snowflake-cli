package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/5HT2/snowflake-cli/util"
	"github.com/bwmarrin/snowflake"
)

var (
	snowflakeInput = flag.Int64("s", 0, "Snowflake to parse")
	formatStr      = flag.String("f", "%u - %p", "Format string, options: %u (unix ms), %p (RFC3339), %t (<t:unix>), %t:S (<t:unix:style>)")
	twitterEpoch   = flag.Bool("twitter", false, "Set to parse Twitter snowflake instead of Discord")
)

func main() {
	flag.Parse()

	if *snowflakeInput == 0 {
		flag.Usage()
		return
	}

	if *twitterEpoch != true {
		snowflake.Epoch = util.DiscordEpoch
	}

	id := snowflake.ParseInt64(*snowflakeInput)
	st := util.SnowflakeTime{Time: time.UnixMilli(id.Time())}
	fmt.Printf("%s\n", st.FormatSnowflake(*formatStr))
}
