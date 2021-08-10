package docker

import (
	"net/http"
	"time"

	"github.com/projecteru2/barrel/cni/subhandler"
	"github.com/projecteru2/barrel/proxy"
	"github.com/projecteru2/barrel/vessel"
)

// NewHandler .
func NewHandler(dockerDaemonSocket string, dialTimeout time.Duration, cniBase *subhandler.Base, vess vessel.Helper) http.Handler {
	client := newHTTPClient(dockerDaemonSocket, dialTimeout)

	inspectAgent := newContainerInspectAgent(client)
	return proxy.HTTPProxyHandler{
		Handlers: []proxy.RequestHandler{
			newContainerCreateHandler(client, vess),
			newContainerDeleteHandler(client, vess, inspectAgent, cniBase),
			newContainerInspectHandler(client, vess),
			newContainerPruneHandle(client, vess),
			newNetworkConnectHandler(client, vess, inspectAgent),
			newNetworkDisconnectHandler(client, vess, inspectAgent),
		},
		HTTPClient: client,
	}
}

// NewSimpleHandler .
func NewSimpleHandler(dockerDaemonSocket string, dialTimeout time.Duration) http.Handler {
	client := newHTTPClient(dockerDaemonSocket, dialTimeout)

	return proxy.HTTPProxyHandler{
		HTTPClient: client,
	}
}
