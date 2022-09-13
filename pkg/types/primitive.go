package types

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// ToMB converts the bytes to MB
func ToMB(bytes int64) int64 {
	return bytes / MB
}

// GetOffsetMB returns missing bytes, which are needed to achieve MB
func GetOffsetMB(bytes int64) int64{
	return bytes % MB
}