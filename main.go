package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/5HT2/snowflake-cli/util"
)

var (
	snowflakeInput = flag.Int64("s", 0, "Snowflake to parse")
	formatStr      = flag.String("f", "%u - %p", "Format string, options: %u (unix ms), %p (RFC3339), %t (<t:unix>), %t:S (<t:unix:style>)")
	twitterEpoch   = flag.Bool("t", false, "Set to parse Twitter snowflake instead of Discord")
)

func main() {
	flag.Parse()

	// If the user didn't set it with -s, try grabbing it from the last arg
	if *snowflakeInput == 0 && len(flag.Args()) > 0 {
		if i, err := strconv.Atoi(flag.Args()[len(flag.Args())-1]); err == nil {
			*snowflakeInput = int64(i)
		}
	}

	// Either the last arg contained a snowflake without -s, or it didn't, either way the user messed up the command
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
