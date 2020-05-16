package as

import (
	"testing"

	as "github.com/aerospike/aerospike-client-go"
)

var hosts = []*as.Host{
	as.NewHost("10.14.41.54", 3000),
	as.NewHost("10.14.41.55", 3000),
	as.NewHost("10.14.41.56", 3000),
}

// Get
func TestGet(t *testing.T) {
	c, err := as.NewClientWithPolicyAndHost(nil, hosts...)
	if err != nil {
		t.Fatalf("client init failed: %v", hosts)
	}
	defer c.Close()
	p := as.NewPolicy()
	p.ReplicaPolicy = as.MASTER
	key1, _ := as.NewKey("dsp", "inverted", "10005")
	r, err := c.Get(p, key1, "adgouplist")
	if err != nil {
		t.Errorf("invoke Get failed: %v", err)
	}
	t.Logf("result: %s", r.Bins["adgouplist"])
}

// BatchGet
func TestBatchGet(t *testing.T) {
	c, err := as.NewClientWithPolicyAndHost(nil, hosts...)
	if err != nil {
		t.Fatalf("client init failed: %v", hosts)
	}
	defer c.Close()
	p := as.NewBatchPolicy()
	p.ReplicaPolicy = as.MASTER
	var key1, _ = as.NewKey("dsp", "inverted", "10005")
	var key2, _ = as.NewKey("dsp", "inverted", "10006")
	var key3, _ = as.NewKey("dsp", "inverted", "10007")
	keys := []*as.Key{key1, key2, key3}
	records, err := c.BatchGet(p, keys, []string{"adgouplist"}...)
	if err != nil {
		t.Errorf("invoke BatchGet failed: %v", err)
	}
	t.Logf("results size: %d", len(records))
}

// Get Benchmark测试
func Benchmark_Get(t *testing.B) {
	t.StopTimer()
	client, _ := as.NewClientWithPolicyAndHost(nil, hosts...)
	defer client.Close()
	p := as.NewPolicy()
	p.ReplicaPolicy = as.MASTER
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		var key1, _ = as.NewKey("dsp", "inverted", "10005")
		client.Get(p, key1, "adgouplist")
	}
}

// BatchGet Benchmark
func Benchmark_BatchGet(t *testing.B) {
	t.StopTimer()
	c, err := as.NewClientWithPolicyAndHost(nil, hosts...)
	if err != nil {
		t.Fatalf("client init failed: %v", hosts)
	}
	defer c.Close()
	p := as.NewBatchPolicy()
	p.ReplicaPolicy = as.MASTER
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		var key1, _ = as.NewKey("dsp", "inverted", "10005")
		var key2, _ = as.NewKey("dsp", "inverted", "10006")
		var key3, _ = as.NewKey("dsp", "inverted", "10007")
		keys := []*as.Key{key1, key2, key3}
		c.BatchGet(p, keys, []string{"adgouplist"}...)
	}
}
