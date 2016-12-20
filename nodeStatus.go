package blackfish

type nodeStatus byte

const (
	STATUS_JOINED nodeStatus = iota
	STATUS_ALIVE
	STATUS_FORWARD_TO
	STATUS_LEFT
	STATUS_DIED
	STATUS_SUSPECTED
)

func (s nodeStatus) String() string {
	switch s {
	case STATUS_JOINED:
		return "JOINED"
	case STATUS_ALIVE:
		return "ALIVE"
	case STATUS_FORWARD_TO:
		return "FORWARD_TO"
	case STATUS_LEFT:
		return "LEFT"
	case STATUS_DIED:
		return "DIED"
	case STATUS_SUSPECTED: // NOT YET IMPLEMENTED
		return "SUSPECTED"
	default:
		return "UNDEFINED"
	}
}
