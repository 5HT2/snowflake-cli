package util

type Snowflake int64

var (
	// IsTwitter toggles which epoch to use
	IsTwitter = false

	// TwitterEpoch is set to the Twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// Docs: https://github.com/twitter-archive/snowflake/blob/b3f6a3c6ca8e1b6847baa6ff42bf72201e2c2231/src/main/scala/com/twitter/service/snowflake/IdWorker.scala#L25
	TwitterEpoch int64 = 1288834974657

	// DiscordEpoch is set to the Discord snowflake epoch of Jan 01 2015 00:00:00 UTC in milliseconds
	// Docs: https://discord.com/developers/docs/reference#snowflakes-snowflake-id-format-structure-left-to-right
	DiscordEpoch int64 = 1420070400000
)

func (t Snowflake) UnixMilli() int64 {
	epoch := DiscordEpoch
	if IsTwitter {
		epoch = TwitterEpoch
	}

	return (int64(t) >> 22) + epoch
}
