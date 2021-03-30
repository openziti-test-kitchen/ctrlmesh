package ctrlmesh

type Config struct {
	ClusterType      string `cf:"cluster_type"`
	BindAddress      string `cf:"bind_address"`
	AdvertiseAddress string `cf:"advertise_address"`
}
