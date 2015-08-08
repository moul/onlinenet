package api

import (
	"strconv"
	"strings"
)

// Servers

type Server struct {
	Identifier int    `json:"id",omitempty`
	Offer      string `json:"offer",omitempty`
	Hostname   string `json:"hostname",omitempty`
	Os         struct {
		Name    string `json:"name",omitempty`
		Version string `json:"version",omitempty`
	} `json:"os",omitempty`
	Version             string `json:"version",omitempty`
	Power               string `json:"power",omitempty`
	BootMode            string `json:"boot_mode",omitempty`
	LastReboot          string `json:"last_reboot",omitempty` // FIXME: Date
	AntiDDOS            bool   `json:"anti_ddos",omitempty`
	HardwareWatch       bool   `json:"hardware_watch",omitempty`
	ProactiveMonitoring bool   `json:"proactive_monitoring",omitempty`
	Support             string `json:"support",omitempty`
	Abuse               string `json:"abuse",omitempty`
	Location            struct {
		Datacenter string `json:"datacenter",omitempty`
		Room       string `json:"room",omitempty`
		Zone       string `json:"zone",omitempty`
		Line       string `json:"line",omitempty`
		Rack       int    `json:"rack",omitempty`
		Block      string `json:"block",omitempty`
		Position   int    `json:"positoin",omitempty`
	} `json:"location",omitempty`
	Network struct {
		Ip      []string `json:"ip",omitempty`
		Private []string `json:"private",omitempty`
		Ipfo    []string `json:"ipfo",omitempty`
	} `json:"network",omitempty`
	Ip []struct {
		Address         string `json:"address",omitempty`
		Type            string `json:"type",omitempty`
		Reverse         string `json:"reverse",omitempty`
		Mac             string `json:"mac",omitempty`
		SwitchPortState string `json:"switch_port_state",omitempty`
		Destination     string `json:"destination",omitempty`
		Server          struct {
			Ref string `json:"$ref",omitempty`
		} `json:"server",omitempty`
		Status string `json:"status",omitempty`
	} `json:"ip",omitempty`
	Contacts struct {
		Owner string `json:"owner",omitempty`
		Tech  string `json:"tech",omitempty`
	} `json:"contacts",omitempty`
	Disks []struct {
		Ref string `json:"$ref",omitempty`
	} `json:"disks",omitempty`
	DriveArrays []struct {
		Disks []struct {
			Ref string `json:"$ref",omitempty`
		} `json:"disks",omitempty`
	} `json:"drive_arrays",omitempty`
	Bmc struct {
		SessionKey *string `json:"session_key",omitempty`
	}
}

type GetServerResp Server

type ServerPath string

type ListServersResp []ServerPath

func (r *ServerPath) Identifier() int {
	idStr := strings.Split(string(*r), "/")[4]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}
	return idInt
}

func (r *ServerPath) Get(c Client) (*GetServerResp, error) {
	return c.GetServer(r.Identifier())
}

// Abuses

type Abuse struct {
	Identifier   int    `json:"id",omitempty`
	Date         string `json:"date",omitempty` // FIXME: Date
	Type         string `json:"type",omitempty`
	Status       string `json:"status",omitempty`
	Service      string `json:"service",omitempty`
	SendDate     string `json:"send_date",omitempty` // FIXME: Date
	Sender       string `json:"senderd",omitempty`
	Description  string `json:"description",omitempty`
	ResolvedDate string `json:"resolved_date",omitempty` // FIXME: Date
	resolver     string `json:"resolver",omitempty`
	answer       string `json:"answer",omitempty`
	solution     string `json:"solution",omitempty`
}

type ListAbusesResp []Abuse

// User

type User struct {
	Identifier int    `json:"id",omitempty`
	Login      string `json:"login",omitempty`
	Email      string `json:"email",omitempty`
	FirstName  string `json:"first_name",omitempty`
	LastName   string `json:"last_name",omitempty`
	Company    string `json:"company",omitempty`
}

type GetCurrentUserResp User
