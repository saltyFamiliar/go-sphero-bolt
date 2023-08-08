package flag

const (
	IsResponse               = 0b1
	RequestResponse          = 0b10
	RequestOnlyErrorResponse = 0b100
	IsActivity               = 0b1000
	HasTargetID              = 0b10000
	HasSourceID              = 0b100000
	Unused                   = 0b1000000
	ExtendedFlags            = 0b10000000
	None                     = 0b0
)
