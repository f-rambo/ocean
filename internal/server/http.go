package server

import (
	"time"

	appv1alpha1 "github.com/f-rambo/ocean/api/app/v1alpha1"
	clusterv1alpha1 "github.com/f-rambo/ocean/api/cluster/v1alpha1"
	projectv1alpha1 "github.com/f-rambo/ocean/api/project/v1alpha1"
	servicev1alpha1 "github.com/f-rambo/ocean/api/service/v1alpha1"
	userv1alpha1 "github.com/f-rambo/ocean/api/user/v1alpha1"
	"github.com/f-rambo/ocean/internal/conf"
	"github.com/f-rambo/ocean/internal/interfaces"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, cluster *interfaces.ClusterInterface, app *interfaces.AppInterface, services *interfaces.ServicesInterface, user *interfaces.UserInterface, project *interfaces.ProjectInterface, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(NewAuthServer(user)).Match(NewWhiteListMatcher()).Build(),
			recovery.Recovery(),
			metadata.Server(),
		),
		http.Timeout(10 * time.Minute),
	}
	cserver := c.Server
	netWork := cserver.HTTP.GetNetwork()
	if netWork != "" {
		opts = append(opts, http.Network(netWork))
	}
	addr := cserver.HTTP.GetAddr()
	if addr != "" {
		opts = append(opts, http.Address(addr))
	}

	opts = append(opts, http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)))

	srv := http.NewServer(opts...)
	clusterv1alpha1.RegisterClusterInterfaceHTTPServer(srv, cluster)
	appv1alpha1.RegisterAppInterfaceHTTPServer(srv, app)
	servicev1alpha1.RegisterServiceInterfaceHTTPServer(srv, services)
	userv1alpha1.RegisterUserInterfaceHTTPServer(srv, user)
	projectv1alpha1.RegisterProjectServiceHTTPServer(srv, project)
	return srv
}
