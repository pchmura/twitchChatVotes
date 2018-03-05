package bot

import (
	"bufio"
	"fmt"
	"github.com/pchmura/twitchChatVotes/Config"
	"net"
	"net/textproto"
	"strings"
	"time"
)

type Bot struct {
	server  string
	port    string
	nick    string
	oauth   string
	channel string
	conn    net.Conn
}

func NewBot(channelName string) *Bot {
	return &Bot{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		nick:    "sentimentBot", //Change to your Twitch username
		oauth:   Config.OAUTH,
		channel: channelName, //Change to your channel
		conn:    nil,
	}
}

func (bot *Bot) Connect() {
	var err error
	fmt.Printf("Connecting to server...\n")
	bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	if err != nil {
		fmt.Printf("Could not connect to Twitch IRC server. Reconnecting in 5 seconds...\n")
		time.Sleep(5 * time.Second)
		bot.Connect()
	}
	fmt.Printf("Connected to IRC server %s\n", bot.server)
}

func RunBot(channelName string, voteOptionA string, voteOptionB string) {
	ircbot := NewBot(channelName)
	ircbot.Connect()
	fmt.Fprintf(ircbot.conn, "USER %s 8 * :%s\r\n", ircbot.nick, ircbot.nick)
	fmt.Fprintf(ircbot.conn, "PASS %s\r\n", ircbot.oauth)
	fmt.Fprintf(ircbot.conn, "NICK %s\r\n", ircbot.nick)
	fmt.Fprintf(ircbot.conn, "JOIN %s\r\n", ircbot.channel)
	defer ircbot.conn.Close()
	reader := bufio.NewReader(ircbot.conn)
	tp := textproto.NewReader(reader)
	strCounter := 0
	for {
		line, err := tp.ReadLine()
		if err != nil {
			break // break loop on errors
		}
		//fmt.Println(line)
		if strings.Contains(line, "PING") {
			pongdata := strings.Split(line, "PING ")
			fmt.Fprintf(ircbot.conn, "PONG %s\r\n", pongdata[1])
		} else if strings.Contains(line, voteOptionA) {
			strCounter++
			fmt.Println(strCounter)
		}
	}

}
