package definitions

type Operation struct {
	// Filled in dynamically prerun
	SrvContainerName  string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	SrvContainerID    string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	DataContainerName string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	DataContainerID   string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	ContainerNumber   int               `json:",omitempty,omitzero" yaml:",omitempty" toml:",omitempty,omitzero"`
	Restart           string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	Remove            bool              `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	Privileged        bool              `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	Attach            bool              `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	AppName           string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	DockerHostConn    string            `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	Labels            map[string]string `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	PublishAllPorts   bool              `json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	CapAdd            []string          `mapstructure:",omitempty", json:",omitempty" yaml:",omitempty" toml:",omitempty"`
	CapDrop           []string          `mapstructure:",omitempty", json:",omitempty" yaml:",omitempty" toml:",omitempty"`
}

func BlankOperation() *Operation {
	return &Operation{}
}
