package app

import (
	"github.com/ssibrahimbas/cli-todos/src/cli"
	"github.com/ssibrahimbas/cli-todos/src/file_stream"
	"github.com/ssibrahimbas/cli-todos/src/internal"
)

type App struct {
	Repo *internal.Repo
	Srv  *internal.Srv
	Fs *file_stream.FileStream
	Cli *cli.Cli
}

func New() *App {
	return &App{}
}

func (a *App) Init() *App {
	a.loadInternal()
	a.loadCli()
	a.loadFileStream()
	return a
}

func (a *App) Start() error {
	return a.Cli.Run()
}

func (a *App) loadFileStream() {
	a.Fs = file_stream.New()
}

func (a *App) loadInternal() {
	a.Repo = internal.NewRepo(&internal.RepoConfig{})
	a.Srv = internal.NewSrv(&internal.SrvConfig{
		Repo: a.Repo,
	})
}

func (a *App) loadCli() {
	a.Cli = cli.New(&cli.CliConfig{
		Repo: a.Repo,
		Srv: a.Srv,
		FileStream: a.Fs,
		FileName: "static/todos.json",
	})
}