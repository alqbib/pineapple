package domain

type UserInfo struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Health struct {
	Epoch                    string `json:"epoch"`
	Timestamp                string `json:"timestamp"`
	Cluster                  string `json:"cluster"`
	Status                   string `json:"status"`
	NodeTotal                string `json:"node.total"`
	NodeData                 string `json:"node.data"`
	DiscoveredClusterManager string `json:"discovered_cluster_manager"`
	Shards                   string `json:"shards"`
	Pri                      string `json:"pri"`
	Relo                     string `json:"relo"`
	Init                     string `json:"init"`
	Unassign                 string `json:"unassign"`
	PendingTasks             string `json:"pending_tasks"`
	MaxTaskWaitTime          string `json:"max_task_wait_time"`
	ActiveShardsPercent      string `json:"active_shards_percent"`
}

type Node struct {
	IP             string `json:"ip"`
	HeapPercent    string `json:"heap.percent"`
	RamPercent     string `json:"ram.percent"`
	Cpu            string `json:"cpu"`
	Load1m         string `json:"load_1m"`
	Load5m         string `json:"load_5m"`
	Load15m        string `json:"load_15m"`
	NodeRole       string `json:"node.role"`
	NodeRoles      string `json:"node.roles"`
	ClusterManager string `json:"cluster_manager"`
	Name           string `json:"name"`
}

type Indices struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	Uuid         string `json:"uuid"`
	Pri          string `json:"pri"`
	Rep          string `json:"rep"`
	DocsCnt      string `json:"docs.count"`
	DocsDel      string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PriStoreSize string `json:"pri.store.size"`
}

type Shards struct {
	Index  string `json:"index"`
	Shard  string `json:"shard"`
	Prirep string `json:"prirep"`
	State  string `json:"state"`
	Docs   string `json:"docs"`
	Store  string `json:"store"`
	IP     string `json:"ip"`
	Node   string `json:"node"`
}
