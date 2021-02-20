package main

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dotvezz/trading-badly/internal/config"
	"github.com/dotvezz/trading-badly/internal/messaging"
	"github.com/dotvezz/trading-badly/internal/messaging/request"
	"image/jpeg"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	discord, err := discordgo.New("Bot " + config.APISecret())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to Discord: %s", err))
	}
	discord.Identify.Intents = discordgo.IntentsGuildMessages + discordgo.IntentsDirectMessages + discordgo.IntentsGuilds

	err = discord.Open()
	defer func() {
		_ = discord.Close()
	}()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	setupMessageHandler(discord)
	listen()
}

func listen() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func setupMessageHandler(session *discordgo.Session) {
	session.AddHandler(handleAll(log.Println))
}

func handleAll(printLog func(...interface{})) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		args := strings.Split(m.Content, " ")

		// Handle private messages to the bot
		c, err := s.Channel(m.ChannelID)
		if err != nil {
			printLog(err)
			return
		}

		req := messaging.Request{
			Author:  m.Author.ID,
			Channel: m.ChannelID,
		}
		resp := &messaging.Response{}

		// Handle direct messages
		if len(c.Recipients) == 1 && len(c.GuildID) > 0 {
			err = request.Handle(req, resp, args...)
			if err != nil {
				printLog(err)
			}
		}

		// Handle channel messages if they ping the bot
		if len(m.Mentions) >= 1 {
			pingedUserID := strings.Trim(args[0], "<@!>")
			if pingedUserID == s.State.User.ID {

				err = request.Handle(req, resp, args[1:]...)
				if err != nil {
					printLog(err)
				}
			}
		}

		responseMessage := &discordgo.MessageSend{Content: resp.Body}
		if resp.Img != nil {
			buffer := &bytes.Buffer{}

			_ = jpeg.Encode(buffer, resp.Img, &jpeg.Options{})
			responseMessage.File = &discordgo.File{
				Name:        "attachment.jpg",
				ContentType: "image/jpg",
				Reader:      buffer,
			}
		}

		if resp.TextFile != nil {
			responseMessage.File = &discordgo.File{
				Name:        "attachment.txt",
				ContentType: "text",
				Reader:      bytes.NewBuffer(resp.TextFile),
			}
		}

		_, err = s.ChannelMessageSendComplex(m.ChannelID, responseMessage)
		if err != nil {
			printLog(err)
		}
	}
}
