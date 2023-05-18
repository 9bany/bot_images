package server

import (
	"context"
	"log"

	"github.com/9bany/bot_workflows/config"
	"github.com/9bany/bot_workflows/server/cmds"
	"github.com/shomali11/slacker"
)

type Server struct {
	slackBot *slacker.Slacker
}

func NewBot(config config.Config) (*Server, error) {

	bot := slacker.NewClient(config.BotToken, config.AppToken)

	return &Server{
		slackBot: bot,
	}, nil

}

func (server *Server) Start() error {

	server.initDefault()
	server.initCommands()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := server.slackBot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (server *Server) initDefault() {
	server.slackBot.Init(func() {
		log.Println("Connected!")
	})

	server.slackBot.Err(func(err string) {
		log.Println(err)
	})

	server.slackBot.DefaultEvent(func(event interface{}) {
		log.Println(event)
	})

}

func (server *Server) initCommands() {

	definition := &slacker.CommandDefinition{
		Description: "help!",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Your own help function...")
		},
	}
	// commands
	server.slackBot.Command(cmds.PingCommandDefinition())
	// help
	server.slackBot.Help(definition)
}
