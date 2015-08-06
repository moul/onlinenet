package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshallServer(t *testing.T) {
	buff := []byte(`{"id":424242,"offer":"Dedibox XXL","hostname":"dedibox-ftw","os":{"name":"ubuntu","version":"14.04_LTS-server"},"power":"ON","boot_mode":"normal","last_reboot":"2014-09-15T11:04:49.000Z","anti_ddos":true,"hardware_watch":true,"proactive_monitoring":true,"support":"Basic service level","abuse":"mail@example.com","location":{"datacenter":"DC3","room":"4","zone":"4-6","line":"C","rack":12,"block":"K","position":4},"network":{"ip":["1.2.3.4"],"private":[],"ipfo":["5.6.7.8"]},"ip":[{"address":"1.2.3.4","type":"public","reverse":"dedibox-ftw.dedibox-fan.fr.","mac":"12:34:56:78:9a:bc","switch_port_state":"up"},{"address":"5.6.7.8","type":"failover","reverse":null,"mac":null,"destination":"1.2.3.4","server":{"$ref":"\/api\/v1\/server\/424242"},"status":"active"}],"contacts":{"owner":"dedibox-fan","tech":"dedibox-fan"},"disks":[{"$ref":"\/api\/v1\/server\/hardware\/disk\/242424"}],"drive_arrays":[{"disks":[{"$ref":"\/api\/v1\/server\/hardware\/disk\/242424"}]}],"bmc":{"session_key":null}}`)

	var server Server
	err := json.Unmarshal(buff, &server)

	assert.Nil(t, err)
	assert.Equal(t, server.Hostname, "dedibox-ftw")
	assert.Equal(t, server.Os.Name, "ubuntu")
}
