package gonduit

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/thought-machine/gonduit/core"
	"github.com/thought-machine/gonduit/responses"
	"github.com/thought-machine/gonduit/util"
)

// A Dialer contains options for connecting to an address.
type Dialer struct {
	ClientName        string
	ClientVersion     string
	ClientDescription string
}

// Dial connects to conduit and confirms the API capabilities for future calls.
func Dial(host string, options *core.ClientOptions) (*Conn, error) {
	var d Dialer

	d.ClientName = "gonduit"
	d.ClientVersion = "1"

	return d.Dial(host, options)
}

// DialFromArcrc connects to conduit. If the API token is not set on the options, it attempts to
// load one from ~/.arcrc.
func DialFromArcrc(host string,	options *core.ClientOptions) (*Conn, error) {
	if options.APIToken != "" {
		return Dial(host, options)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("Cannot determine home dir: %s", err)
	}
	f, err := os.Open(path.Join(home, ".arcrc"))
	if err != nil {
		return nil, fmt.Errorf("Cannot open ~/.arcrc: %s", err)
	}
	defer f.Close()
	arcrc := struct {
		Hosts map[string]struct {
			Token string `json:"token"`
		} `json:"hosts"`
	}{}
	if err := json.NewDecoder(f).Decode(&arcrc); err != nil {
		return nil, fmt.Errorf("Failed to decode ~/.arcrc: %s", err)
	}
	for k, v := range arcrc.Hosts {
		if strings.Contains(k, host) {
			options.APIToken = v.Token
			return Dial(host, options)
		}
	}
	return nil, fmt.Errorf("Failed to find appropriate token in .arcrc")
}

// Dial connects to conduit and confirms the API capabilities for future calls.
func (d *Dialer) Dial(
	host string,
	options *core.ClientOptions,
) (*Conn, error) {
	var res responses.ConduitCapabilitiesResponse

	// We use conduit.connect for authentication and it establishes a session.
	err := core.PerformCall(
		core.GetEndpointURI(host, "conduit.getcapabilities"),
		nil,
		&res,
		options,
	)
	if err != nil {
		return nil, err
	}

	// Now, we need to assert that the conduit API supports this client.
	assertSupportedCapabilities(res, options)

	conn := Conn{
		host:         host,
		capabilities: &res,
		dialer:       d,
		options:      options,
	}

	return &conn, nil
}

func assertSupportedCapabilities(
	res responses.ConduitCapabilitiesResponse,
	options *core.ClientOptions,
) error {
	if options.APIToken != "" {
		if !util.ContainsString(res.Authentication, "token") {
			return core.ErrTokenAuthUnsupported
		}
	}

	if options.Cert != "" {
		if !util.ContainsString(res.Authentication, "session") {
			return core.ErrSessionAuthUnsupported
		}
	}

	if !util.ContainsString(res.Input, "urlencoded") {
		return core.ErrURLEncodedInputUnsupported
	}

	if !util.ContainsString(res.Output, "json") {
		return core.ErrJSONOutputUnsupported
	}

	return nil
}
