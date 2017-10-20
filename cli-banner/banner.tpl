{{ .Cli.BannerColor }}
              v .   ._, |_  .,
           `-._\/  .  \ /    |/_t
               \\  _\, y | \//
         _\_.___\\, \\/ -.\||
           `7-,--.`._||  / / ,r
           /'     `-. `./ / |/_.'
              3      |    |//
                     |_   e/
                     |-   |
                     |   =|
                     |    |
--------------------/ ,  . \--------._{{.Cli.ImportantColor}}
    project:    {{.Cli.Title}}
    version:    {{.Cli.Version}}
    statement:  {{.Cli.Statement}}{{.Cli.InfoColor}}
    platform:   {{.GOOS}}_{{.GOARCH}} {{.GoVersion}}{{.Cli.WarningColor}}
>>>     launched on {{ .Now "Monday, 2 Jan 2006" }}     <<<

