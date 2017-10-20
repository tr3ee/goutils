package banner

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	colorable "github.com/mattn/go-colorable"
)

type vars struct {
	GoVersion string
	GOOS      string
	GOARCH    string
	Cli       cli
}

func (v vars) Now(layout string) string {
	return time.Now().Format(layout)
}

// DefaultBannerFile is the default file to render
const DefaultBannerFile = "banner.tpl"

func Autoload(title, version, statement string) {
	out := colorable.NewColorableStdout()
	c := cli{
		Title:     title,
		Version:   version,
		Statement: statement,
	}
	// restore default color
	defer func() {
		out.Write([]byte(c.DefaultColor()))
	}()
	// panic recover
	defer func() {
		if err := recover(); err != nil {
			out.Write([]byte(c.Error("\n[error] Banner-autoload:" + err.(error).Error())))
		}
	}()
	// get this file path
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("Unable to access the current file path"))
	}
	path := filepath.Dir(file)
	bytes, err := ioutil.ReadFile(path + "/" + DefaultBannerFile)
	if err != nil {
		panic(fmt.Errorf("Unable to access the render file \"" + DefaultBannerFile + "\""))
	}
	// render template
	if err := ColorableRender(out, string(bytes), c); err != nil {
		panic(err)
	}
}

// ColorableRender help user to render the colorable-template file
func ColorableRender(out io.Writer, content string, cmdline cli) error {
	t, err := template.New("banner").Parse(content)
	if err != nil {
		return err
	}
	return t.Execute(out, vars{
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		cmdline,
	})
}
