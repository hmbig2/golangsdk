package nics

type NicsResp struct {
	InterfaceAttachments []Nic                  `json:"interfaceAttachments"`
	AttachableQuantity   AttachableQuantityInfo `json:"attachableQuantity"`
}

type FixedIP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

type Nic struct {
	PortState           string    `json:"port_state"`
	FixedIps            []FixedIP `json:"fixed_ips"`
	NetworkId           string    `json:"net_id"`
	PortId              string    `json:"port_id"`
	MacAddr             string    `json:"mac_addr"`
	DeleteOnTermination bool      `json:"delete_on_termination"`
	// options: virtio, hinic.  Default is virtio。
	DriverMode    string `json:"driver_mode"`
	MinRate       int    `json:"min_rate"`
	MultiqueueNum int    `json:"multiqueue_num"`
	PciAddress    string `json:"pci_address"`
}

type AttachableQuantityInfo struct {
	FreeNic int `json:"free_nic"`
}
