package ctrlmesh

type Config struct {
	BindAddress      string `cf:"bind_address"`
	AdvertiseAddress string `cf:"advertise_address"`
	InitialPeer      string `cf:"initial_peer"`
}
