package main

import (
	"context"
	"fmt"
	"log"
	"naptalie-api/api/routes"
	discordclient "naptalie-api/discord-client"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func startHTTPServer(ctx context.Context) error {
	http.HandleFunc("/weather", routes.HandleDiscordWebhookWeather)

	server := &http.Server{
		Addr: ":8090",
	}

	// yay goroutines
	go func() {
		log.Println("api server started on :8090")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Oh noer Cleor I touched the wata :( :%v", err)
		}
	}()

	<-ctx.Done()

	// Graceful shutdown with timeout
	log.Println("HTTP server shutting down...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	return server.Shutdown(shutdownCtx)
}

func startDiscordBot(ctx context.Context, token string) error {
	// Create a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("error creating Discord session: %w", err)
	}

	// Register event handlers
	dg.AddHandler(discordclient.MessageCreate)
	dg.AddHandler(discordclient.Ready)

	// Set intents (adjust based on your needs)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	// Open a websocket connection to Discord
	err = dg.Open()
	if err != nil {
		return fmt.Errorf("error opening Discord connection: %w", err)
	}
	defer dg.Close()

	log.Println("Discord bot is running...")

	// Wait for context cancellation
	<-ctx.Done()
	log.Println("Discord bot shutting down...")
	return nil
}

func main() {
	discordToken := os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("DISCORD_TOKEN must be set")
	}

	//create waitgroup
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startDiscordBot(ctx, discordToken); err != nil {
			log.Printf("discord bot error: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := startHTTPServer(ctx); err != nil {
			log.Printf("http server started: %v", err)
		}
	}()
	// wait for interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down!!!!!!")

	// cancel context to signal shutdown
	cancel()

	wg.Wait()
	log.Println("Shutdown complete!!!!")
}
