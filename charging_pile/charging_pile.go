package charging_pile

type ChargingPile struct {
	TypeID                uint8
	RatedPower            float32
	ElectricCurrentType   uint8
	VoltageInput          uint32
	VoltageOutput         uint32
	ElectricCurrentOutput uint32
	GunNum                uint32
	AmmeterNum            float32
}
