package open

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/opensearch-project/opensearch-go"

	"opensearch-tool/tool/domain"
	"opensearch-tool/tool/util"
)

type OpenTool struct {
	client *opensearch.Client
}

func NewOpenTool(userInfo domain.UserInfo) (OpenTool, error) {
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{userInfo.Host},
		Username:  userInfo.User, // For testing only. Don't store credentials in code.
		Password:  userInfo.Password,
	})
	return OpenTool{
		client: client,
	}, err
}

func (tool *OpenTool) Info() (map[string]interface{}, error) {
	info, err := tool.client.Info()
	if err != nil {
		return nil, err
	}
	s := util.TrimStr(info.String())
	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(s), &m)
	return m, err
}

func (tool *OpenTool) CatHealth() ([]domain.Health, error) {
	results := make([]domain.Health, 0)
	health, err := tool.client.Cat.Health()
	if err != nil {
		return nil, err
	}
	defer health.Body.Close()
	r := regexp.MustCompile("[^\\s]+")
	lines := util.StrToArray(health.String())
	for _, line := range lines {
		var info domain.Health
		res := r.FindAllString(line, -1)
		info.Epoch = res[0]
		info.Timestamp = res[1]
		info.Cluster = res[2]
		info.Status = res[3]
		info.NodeTotal = res[4]
		info.NodeData = res[5]
		info.DiscoveredClusterManager = res[6]
		info.Shards = res[7]
		info.Pri = res[8]
		info.Relo = res[9]
		info.Init = res[10]
		info.Unassign = res[11]
		info.PendingTasks = res[12]
		info.MaxTaskWaitTime = res[13]
		info.ActiveShardsPercent = res[14]
		results = append(results, info)
	}
	return results, nil
}

func (tool *OpenTool) CatNodes() ([]domain.Node, error) {
	results := make([]domain.Node, 0)
	nodes, err := tool.client.Cat.Nodes()
	if err != nil {
		return nil, err
	}
	defer nodes.Body.Close()
	r := regexp.MustCompile("[^\\s]+")
	lines := util.StrToArray(nodes.String())
	for _, line := range lines {
		var info domain.Node
		res := r.FindAllString(line, -1)
		info.IP = res[0]
		info.HeapPercent = res[1]
		info.RamPercent = res[2]
		info.Cpu = res[3]
		info.Load1m = res[4]
		info.Load5m = res[5]
		info.Load15m = res[6]
		info.NodeRole = res[7]
		info.NodeRoles = res[8]
		info.ClusterManager = res[9]
		info.Name = res[10]
		results = append(results, info)
	}
	return results, nil
}

func (tool *OpenTool) CatIndices() ([]domain.Indices, error) {
	results := make([]domain.Indices, 0)
	indices, err := tool.client.Cat.Indices()
	if err != nil {
		return nil, err
	}
	defer indices.Body.Close()
	r := regexp.MustCompile("[^\\s]+")
	lines := util.StrToArray(indices.String())
	for _, line := range lines {
		var info domain.Indices
		res := r.FindAllString(line, -1)
		info.Health = res[0]
		info.Status = res[1]
		info.Index = res[2]
		info.Uuid = res[3]
		info.Pri = res[4]
		info.Rep = res[5]
		info.DocsCnt = res[6]
		info.DocsDel = res[7]
		info.StoreSize = res[8]
		info.PriStoreSize = res[9]
		results = append(results, info)
	}
	return results, nil
}

func (tool *OpenTool) CatShards() ([]domain.Shards, error) {
	results := make([]domain.Shards, 0)
	shards, err := tool.client.Cat.Shards()
	if err != nil {
		return nil, err
	}
	defer shards.Body.Close()
	r := regexp.MustCompile("[^\\s]+")
	lines := util.StrToArray(shards.String())
	for _, line := range lines {
		var info domain.Shards
		res := r.FindAllString(line, -1)
		info.Index = res[0]
		info.Shard = res[1]
		info.Prirep = res[2]
		info.State = res[3]
		if len(res) == 8 {
			info.Docs = res[4]
			info.Store = res[5]
			info.IP = res[6]
			info.Node = res[7]
		}
		results = append(results, info)
	}
	return results, nil
}
