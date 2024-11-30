package main

import (
	"fmt"
	"html/template"
	"io"
	"src/common/setting"
	custommw "src/middleware"
	"src/route"
	"src/util/dbutil"
	"src/util/routeutil"

	sentry "github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "src/docs"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// CustomValidator implements the echo.Validator interface
type customValidator struct {
	Validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

//	@title			Simple PM API
//	@version		1.0
//	@description	Simple PM API document.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		simplepm.io.io
// @BasePath	/api/v1
func main() {
	dbutil.InitDb()
	e := echo.New()
	e.Debug = true
	e.Validator = &customValidator{Validator: validator.New()}
	t := &Template{
		templates: template.Must(template.ParseGlob("/code/public/views/*.html")),
	}
	e.Renderer = t
	if !setting.DEBUG {
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())
		// sentry setup
		if err := sentry.Init(sentry.ClientOptions{
			Dsn: setting.SENTRY_DSN,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for tracing.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v\n", err)
		}
		e.Use(sentryecho.New(sentryecho.Options{}))
	}
	e.Use(custommw.LangMiddleware)

	apiGroup := e.Group("/api/v1")
	_, pemMap := route.CollectRoutes(apiGroup)
	routeutil.SetPemMap(&pemMap)
	e.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start("0.0.0.0:4000"))
}
