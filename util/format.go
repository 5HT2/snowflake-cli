package util

import (
	"fmt"
	"strings"
	"time"
)

type SnowflakeTime struct {
	time.Time
}

func (t SnowflakeTime) FormatUnix() string {
	return fmt.Sprintf("%v", t.UnixMilli())
}

func (t SnowflakeTime) FormatPretty() string {
	t.Location()
	return t.Format(time.RFC3339)
}

func (t SnowflakeTime) FormatTimestamp(f DiscordFormat) string {
	return fmt.Sprintf("<t:%v%s>", t.Unix(), f)
}

func (t SnowflakeTime) FormatSnowflake(s string) string {
	s = strings.ReplaceAll(s, "%%", "{%}")
	s = strings.ReplaceAll(s, "%u", t.FormatUnix())
	s = strings.ReplaceAll(s, "%p", t.FormatPretty())
	s = strings.ReplaceAll(s, "%t:t", t.FormatTimestamp(DiscordFormatShortTime))
	s = strings.ReplaceAll(s, "%t:T", t.FormatTimestamp(DiscordFormatLongTime))
	s = strings.ReplaceAll(s, "%t:d", t.FormatTimestamp(DiscordFormatShortDate))
	s = strings.ReplaceAll(s, "%t:D", t.FormatTimestamp(DiscordFormatLongDate))
	s = strings.ReplaceAll(s, "%t:f", t.FormatTimestamp(DiscordFormatShortDateTime))
	s = strings.ReplaceAll(s, "%t:F", t.FormatTimestamp(DiscordFormatLongDateTime))
	s = strings.ReplaceAll(s, "%t:R", t.FormatTimestamp(DiscordFormatRelative))
	s = strings.ReplaceAll(s, "%t", t.FormatTimestamp(DiscordFormatDefault))
	return strings.ReplaceAll(s, "{%}", "%")
}
