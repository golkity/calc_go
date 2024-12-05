package application

import (
	"encoding/json"
	"fmt"
	"github.com/yourusername/yourproject/pkg/calc" // Импортируем калькулятор
	"net/http"
	"time"
)

// Response структура для ответа в формате JSON.
type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

// Application — структура приложения.
type Application struct {
	timeout time.Duration
}

// New создает новый экземпляр приложения с указанным таймаутом.
func New(timeout time.Duration) *Application {
	return &Application{timeout: timeout}
}

// calculateHandler обработчик для вычислений.
func (app *Application) calculateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем выражение из параметров запроса
	expression := r.URL.Query().Get("expression")
	if expression == "" {
		http.Error(w, "Expression is required", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для вычислений
	result, err := calc.Calc(expression)

	// Формируем ответ
	resp := Response{}
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Result = result
	}

	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Отправляем ответ в формате JSON
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

// Run запускает HTTP сервер с таймаутом.
func (app *Application) Run() {
	// Обработчик для маршрута /calculate
	http.HandleFunc("/calculate", app.calculateHandler)

	// Создаем сервер с таймаутом
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  app.timeout,
		WriteTimeout: app.timeout,
	}

	// Запуск HTTP-сервера
	fmt.Printf("Server is running on http://localhost:8080\n")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
