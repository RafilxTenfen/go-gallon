package gallons

// Bottle struct to store volume
type Bottle struct {
	VolumeML uint32
}

// NewBottle returns a new Bottle struct
func NewBottle(volume uint32) Bottle {
	return Bottle{
		VolumeML: volume,
	}
}
