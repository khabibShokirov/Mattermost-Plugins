package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/mattermost/mattermost/server/public/pluginapi" // ДОБАВЬТЕ ЭТОТ ИМПОРТ
)

// apiHandler содержит клиент API для использования в HTTP-методах.
type apiHandler struct {
	clientAPI *pluginapi.Client // ЭТО НОВЫЙ КЛИЕНТ API
}

// ServeHTTP обрабатывает HTTP-запросы к плагину.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	// Инициализируем обработчик HTTP-запросов, передавая ему клиент API плагина.
	h := &apiHandler{
		clientAPI: p.client, // Используем p.client, инициализированный в plugin.go
	}

	router := mux.NewRouter()

	// Middleware для требования авторизации пользователя.
	// Теперь используем middleware из apiHandler.
	router.Use(h.MattermostAuthorizationRequired)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Регистрируем хендлер для "/hello", используя метод apiHandler.
	apiRouter.HandleFunc("/hello", h.HelloWorld).Methods(http.MethodGet)

	router.ServeHTTP(w, r)
}

// MattermostAuthorizationRequired является middleware для проверки авторизации Mattermost.
func (a *apiHandler) MattermostAuthorizationRequired(next http.Handler) http.Handler { // Изменили p на a
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("Mattermost-User-ID")
		if userID == "" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// HelloWorld демонстрирует простую конечную точку HTTP.
func (a *apiHandler) HelloWorld(w http.ResponseWriter, r *http.Request) { // Изменили p на a
	if _, err := w.Write([]byte("Hello, world!")); err != nil {
		a.clientAPI.Log.Error("Failed to write response", "error", err) // ИСПРАВЛЕНИЕ: Убрали лишнее ".API"
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
