package models

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type HandlerInterface interface {
	AuthMiddleware(next http.Handler) http.Handler

	// IndexHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	K8SConfigHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, req *http.Request)
	// DashboardHandler(ctx context.Context, w http.ResponseWriter, req *http.Request)
	UserHandler(w http.ResponseWriter, r *http.Request)
	LoadTestHandler(w http.ResponseWriter, req *http.Request)
	FetchResultsHandler(w http.ResponseWriter, req *http.Request)
	MeshOpsHandler(w http.ResponseWriter, req *http.Request)

	GrafanaConfigHandler(w http.ResponseWriter, req *http.Request)
	GrafanaBoardsHandler(w http.ResponseWriter, req *http.Request)
	GrafanaQueryHandler(w http.ResponseWriter, req *http.Request)
}

type HandlerConfig struct {
	SessionName   string
	RefCookieName string
	// SessionUserKey
	// SessionSecret string

	SessionStore sessions.Store

	SaaSTokenName string
	SaaSBaseURL   string

	ByPassAuth bool

	FortioURL string
}

//
// sessionUserKey       = "twitterID"
// 	sessionUserName      = "twitterUserName"
// cookieSuffix         = "_referrer"
// saasTokenName        = "meshery_saas"
