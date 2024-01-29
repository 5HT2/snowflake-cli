package main

import (
	"flag"
	"fmt"

	"github.com/5HT2/snowflake-cli/util"
)

var (
	snowflakeInput = flag.Int64("s", 0, "Snowflake to parse")
	formatStr      = flag.String("f", "%u - %p", "Format string, options: %u (unix ms), %p (RFC3339), %t (<t:unix>), %t:S (<t:unix:style>)")
	twitterEpoch   = flag.Bool("t", false, "Set to parse Twitter snowflake instead of Discord")
)

func main() {
	flag.Parse()

	if *snowflakeInput == 0 {
		flag.Usage()
		return
	}

	if *twitterEpoch {
		util.IsTwitter = true
	}

	s := util.NewSnowflake(*snowflakeInput)
	fmt.Printf("%s\n", s.FormatSnowflake(*formatStr))
}
