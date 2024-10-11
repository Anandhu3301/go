package packet



func Packet12(val int) string {
	if val < 50 {
		return "Its low"
	} else {
		return "Its high"

	}
}