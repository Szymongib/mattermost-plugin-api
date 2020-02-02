package cluster_test

import (
	"github.com/lieut-data/mattermost-plugin-api/cluster"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func ExampleMutex() {
	// Use p.API from your plugin instead.
	pluginAPI := plugin.API(nil)

	m := cluster.NewMutex("key", pluginAPI)
	m.Lock()
	// critical section
	m.Unlock()
}