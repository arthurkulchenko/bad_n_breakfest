package helpers

import(
	"net/http"
	"github.com/arthurkulchenko/bed_n_breakfest/internal/config"
	"fmt"
	"runtime/debug"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClienError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
