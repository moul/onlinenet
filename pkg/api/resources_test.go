package api

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
)

func ExampleDdos_json() []byte {
	return []byte(`
{
    "id": 12345,
    "target": "1.2.3.4",
    "start": "2013-10-24T21:46:39.000Z",
    "end": "2013-10-24T21:55:49.000Z",
    "mitigation": "root",
    "type": "DoS Fragmentation Misuse",
    "max_pps": 261758,
    "max_bps": 2652170368,
    "timeline": [
      {
        "timestamp": 1382651259,
        "pps": 174463,
        "bps": 1780643680
      },
      {
        "timestamp": 1382651319,
        "pps": 261758,
        "bps": 2652170368
      },
      {
        "timestamp": 1382651379,
        "pps": 242217,
        "bps": 2442621952
      },
      {
        "timestamp": 1382651439,
        "pps": 75825,
        "bps": 765489520
      }
    ]
}
`)
}

func ExampleServer_json() []byte {
	return []byte(`
{
  "id": 424242,
  "offer": "Dedibox XXL",
  "hostname": "dedibox-ftw",
  "os": {
    "name": "ubuntu",
    "version": "14.04_LTS-server"
  },
  "power": "ON",
  "boot_mode": "normal",
  "last_reboot": "2014-09-15T11:04:49.000Z",
  "anti_ddos": true,
  "hardware_watch": true,
  "proactive_monitoring": true,
  "support": "Basic service level",
  "abuse": "mail@example.com",
  "location": {
    "datacenter": "DC3",
    "room": "4",
    "zone": "4-6",
    "line": "C",
    "rack": 12,
    "block": "K",
    "position": 4
  },
  "network": {
    "ip": [
      "1.2.3.4"
    ],
    "private": [],
    "ipfo": [
      "5.6.7.8"
    ]
  },
  "ip": [
    {
      "address": "1.2.3.4",
      "type": "public",
      "reverse": "dedibox-ftw.dedibox-fan.fr.",
      "mac": "12:34:56:78:9a:bc",
      "switch_port_state": "up"
    },
    {
      "address": "5.6.7.8",
      "type": "failover",
      "reverse": null,
      "mac": null,
      "destination": "1.2.3.4",
      "server": {
        "$ref": "/api/v1/server/424242"
      },
      "status": "active"
    }
  ],
  "contacts": {
    "owner": "dedibox-fan",
    "tech": "dedibox-fan"
  },
  "disks": [
    {
      "$ref": "/api/v1/server/hardware/disk/242424"
    }
  ],
  "drive_arrays": [
    {
      "disks": [
        {
          "$ref": "/api/v1/server/hardware/disk/242424"
        }
      ]
    }
  ],
  "bmc": {
    "session_key": null
  }
}
`)
}

func TestUnmarshallServer(t *testing.T) {
	buff := ExampleServer_json()

	var server Server
	err := json.Unmarshal(buff, &server)
	assert.Nil(t, err)

	isValid, err := govalidator.ValidateStruct(server)
	assert.Nil(t, err)
	assert.True(t, isValid)

	assert.Equal(t, server.Identifier, 424242)
	assert.Equal(t, server.Offer, "Dedibox XXL")
	assert.Equal(t, server.Hostname, "dedibox-ftw")
	assert.Equal(t, server.Os.Name, "ubuntu")
	assert.Equal(t, server.Os.Version, "14.04_LTS-server")
	assert.Equal(t, server.Power, "ON")
	assert.Equal(t, server.BootMode, "normal")

	lastReboot, err := time.Parse(time.RFC3339, "2014-09-15T11:04:49.000Z")
	assert.Nil(t, err)
	assert.Equal(t, server.LastReboot, lastReboot)

	assert.True(t, server.AntiDDOS)
	assert.True(t, server.HardwareWatch)
	assert.True(t, server.ProactiveMonitoring)
	assert.Equal(t, server.Support, "Basic service level")
	assert.Equal(t, server.Abuse, "mail@example.com")
	assert.Equal(t, server.Location.Datacenter, "DC3")
	assert.Equal(t, server.Location.Room, "4")
	assert.Equal(t, server.Location.Zone, "4-6")
	assert.Equal(t, server.Location.Line, "C")
	assert.Equal(t, server.Location.Rack, 12)
	assert.Equal(t, server.Location.Block, "K")
	assert.Equal(t, server.Location.Position, 4)
	assert.Equal(t, len(server.Network.Ip), 1)
	assert.Equal(t, len(server.Network.Private), 0)
	assert.Equal(t, len(server.Network.Ipfo), 1)
	assert.Equal(t, server.Network.Ip[0], "1.2.3.4")
	assert.Equal(t, server.Network.Ipfo[0], "5.6.7.8")
	assert.Equal(t, len(server.Ip), 2)
	assert.Equal(t, server.Ip[0].Address, "1.2.3.4")
	assert.Equal(t, server.Ip[0].Type, "public")
	assert.Equal(t, server.Ip[0].Reverse, "dedibox-ftw.dedibox-fan.fr.")
	assert.Equal(t, server.Ip[0].Mac, "12:34:56:78:9a:bc")
	assert.Equal(t, server.Ip[0].SwitchPortState, "up")
	assert.Equal(t, server.Ip[1].Address, "5.6.7.8")
	assert.Equal(t, server.Ip[1].Type, "failover")
	assert.Equal(t, server.Ip[1].Reverse, "")
	assert.Equal(t, server.Ip[1].Mac, "")
	assert.Equal(t, server.Ip[1].Destination, "1.2.3.4")
	assert.Equal(t, server.Ip[1].Server.Ref, "/api/v1/server/424242")
	assert.Equal(t, server.Ip[1].Status, "active")
	assert.Equal(t, server.Contacts.Owner, "dedibox-fan")
	assert.Equal(t, server.Contacts.Tech, "dedibox-fan")
	assert.Equal(t, len(server.Disks), 1)
	assert.Equal(t, server.Disks[0].Ref, "/api/v1/server/hardware/disk/242424")
	assert.Equal(t, len(server.DriveArrays), 1)
	assert.Equal(t, len(server.DriveArrays[0].Disks), 1)
	assert.Equal(t, server.DriveArrays[0].Disks[0].Ref, "/api/v1/server/hardware/disk/242424")
	assert.Equal(t, server.Bmc.SessionKey, "")
}

func TestUnmarshallDdos(t *testing.T) {
	buff := ExampleDdos_json()

	var ddos Ddos
	err := json.Unmarshal(buff, &ddos)
	assert.Nil(t, err)

	isValid, err := govalidator.ValidateStruct(ddos)
	assert.Nil(t, err)
	assert.True(t, isValid)

	assert.Equal(t, ddos.Identifier, 12345)
	assert.Equal(t, ddos.Target, "1.2.3.4")

	ddosStart, err := time.Parse(time.RFC3339, "2013-10-24T21:46:39.000Z")
	assert.Nil(t, err)
	assert.Equal(t, ddos.Start, ddosStart)

	ddosEnd, err := time.Parse(time.RFC3339, "2013-10-24T21:55:49.000Z")
	assert.Nil(t, err)
	assert.Equal(t, ddos.End, ddosEnd)

	assert.Equal(t, ddos.Mitigation, "root")
	assert.Equal(t, ddos.MaxPPS, 261758)
	assert.Equal(t, ddos.MaxBPS, 2652170368)
	assert.Equal(t, len(ddos.Timeline), 4)
	assert.Equal(t, ddos.Timeline[0].Timestamp, 1382651259)
	assert.Equal(t, ddos.Timeline[0].PPS, 174463)
	assert.Equal(t, ddos.Timeline[0].BPS, 1780643680)
}

func TestUnmarshallUser(t *testing.T) {
	buff := []byte(`{"id":123456,"login":"johndoe42","email":"technical@example.com","first_name":"John","last_name":"Doe","company":null}`)

	var user User
	err := json.Unmarshal(buff, &user)

	assert.Nil(t, err)
	assert.Equal(t, user.Identifier, 123456)
	assert.Equal(t, user.Login, "johndoe42")
	assert.Equal(t, user.Email, "technical@example.com")
	assert.Equal(t, user.FirstName, "John")
	assert.Equal(t, user.LastName, "Doe")
	assert.Equal(t, user.Company, "")
}
