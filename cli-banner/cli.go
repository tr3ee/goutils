package banner

import "time"

type cli struct {
	Title          string
	Version        string
	Statement      string
	AnsiColor      ansiColor
	AnsiBackground ansiBackground
}

// ------------ color group -----------------

func (c cli) DefaultColor() string {
	return c.AnsiColor.Default()
}

func (c cli) TextColor() string {
	return c.AnsiColor.White()
}

func (c cli) InfoColor() string {
	return c.AnsiColor.Cyan()
}

func (c cli) DebugColor() string {
	return c.AnsiColor.Blue()
}

func (c cli) WarningColor() string {
	return c.AnsiColor.Yellow()
}

func (c cli) SuccessColor() string {
	return c.AnsiColor.Green()
}

func (c cli) ImportantColor() string {
	return c.AnsiColor.Red()
}

func (c cli) ErrorColor() string {
	return c.AnsiColor.BrightRed()
}

// ------------ message group -----------------------

func (c cli) Default(msg string) string {
	return c.DefaultColor() + msg
}

func (c cli) Text(msg string) string {
	return c.TextColor() + msg
}

func (c cli) Info(msg string) string {
	return c.InfoColor() + msg
}

func (c cli) Debug(msg string) string {
	return c.DebugColor() + msg
}

func (c cli) Warning(msg string) string {
	return c.WarningColor() + msg
}

func (c cli) Success(msg string) string {
	return c.SuccessColor() + msg
}

func (c cli) Important(msg string) string {
	return c.ImportantColor() + msg
}

func (c cli) Error(msg string) string {
	return c.ErrorColor() + msg
}

// ------------ cli group -----------------

func (c cli) BannerColor() string {
	switch int(time.Now().Month()) / 3 {
	case 1:
		return c.AnsiColor.BrightGreen()
	case 2:
		return c.AnsiColor.BrightCyan()
	case 3:
		return c.AnsiColor.BrightYellow()
	case 0:
		return c.AnsiColor.BrightWhite()
	default:
		return c.AnsiColor.Default()
	}
}
