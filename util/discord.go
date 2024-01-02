package util

const (
	// DiscordEpoch is set to the Discord snowflake epoch of Jan 01 2015 00:00:00 UTC in milliseconds
	// Docs: https://discord.com/developers/docs/reference#snowflakes-snowflake-id-format-structure-left-to-right
	DiscordEpoch = 1420070400000
)

// DiscordFormat defaults to DiscordFormatShortDateTime
// Docs: https://discord.com/developers/docs/reference#message-formatting-timestamp-styles
type DiscordFormat int64

const (
	DiscordFormatDefault DiscordFormat = iota
	DiscordFormatShortTime
	DiscordFormatLongTime
	DiscordFormatShortDate
	DiscordFormatLongDate
	DiscordFormatShortDateTime
	DiscordFormatLongDateTime
	DiscordFormatRelative
)

func (f DiscordFormat) String() string {
	switch f {
	case DiscordFormatShortTime:
		return ":t"
	case DiscordFormatLongTime:
		return ":T"
	case DiscordFormatShortDate:
		return ":d"
	case DiscordFormatLongDate:
		return ":D"
	case DiscordFormatShortDateTime:
		return ":f"
	case DiscordFormatLongDateTime:
		return ":F"
	case DiscordFormatRelative:
		return ":R"
	default:
		return ""
	}
}
