package main

import (
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	mn "github.com/mikejk8s/gmud/pkg/menus"
	"github.com/mikejk8s/gmud/pkg/models"
	sqlpkg "github.com/mikejk8s/gmud/pkg/mysqlpkg"
	"github.com/mikejk8s/gmud/pkg/tcpserver"
	"github.com/muesli/termenv"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

const (
	host = "localhost"
	port = 2222
)

func pkHandler(ctx ssh.Context, key ssh.PublicKey) bool {
	return true
}
func passHandler(ctx ssh.Context, password string) bool {
	usersConn := sqlpkg.SqlConn{}
	err := usersConn.GetSQLConn("users")
	if err != nil {
		log.Fatalln(err)
	}
	// data := usersDB.Exec(fmt.Sprintf("SELECT password from users.users where username = '%s'", ctx.User())).First(&p)
	user := models.User{}
	rows, _ := usersConn.DB.Query(fmt.Sprintf("SELECT password, email, name, username from users.users where username = '%s'", ctx.User()))
	for rows.Next() {
		err := rows.Scan(&user.Password, &user.Email, &user.Name, &user.Username)
		if err != nil {
			panic(err)
		}
	}
	usersConn.CloseConn()
	if user.Password == password {
		credentialError := user.CheckPassword(password)
		if credentialError != nil {
			return false
		} else {
			return true
		}

	} else {
		return false
	}
}
func main() {
	// Create users schema and users table, migrate if possible.
	go sqlpkg.Migration()
	// Connect to mariadb database and create characters schema + character tables if they don't exist
	initialTableCreation := sqlpkg.SqlConn{}
	err := initialTableCreation.GetSQLConn("characters")
	if err != nil {
		log.Println(err)
	}
	go func() {
		err := initialTableCreation.CreateCharacterTable()
		if err != nil {
			log.Println(err)
		} else {
			initialTableCreation.CloseConn()
		}
	}()
	go func() {
		initialUsersCreation := sqlpkg.SqlConn{}
		err := initialUsersCreation.GetSQLConn("")
		if err != nil {
			log.Fatalln(err)
		}
		initialUsersCreation.CreateUsersTable()
		initialUsersCreation.CloseConn()
	}()
	// Create a new TCP server
	newTCP := tcpserver.TCPServer{}
	newTCP.Port = os.Getenv("TCP_HOST")
	newTCP.Host = os.Getenv("TCP_PORT")
	go newTCP.CreateListener()
	// SSH server begin
	s, err := wish.NewServer(
		ssh.PasswordAuth(passHandler),
		ssh.PublicKeyAuth(pkHandler),
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			lm.Middleware(),
			loginBubbleteaMiddleware(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	s.ConnectionFailedCallback = func(conn net.Conn, err error) {
		log.Println("Connection failed:", err)
	}
	done := make(chan os.Signal, 0)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		err = s.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	// start web based terminal
	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func loginBubbleteaMiddleware() wish.Middleware {
	login := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
		p := tea.NewProgram(m, opts...)
		go func() {
			// if err := p.Start(); err != nil {
			// 	log.Fatalln(err)
			// }
			for {
				<-time.After(1 * time.Second)
				p.Send(timeMsg(time.Now()))
			}
		}()
		return p
	}
	teaHandler := func(s ssh.Session) *tea.Program {
		pty, _, _ := s.Pty()
		m := model{
			SSHSession: s,
			Width:      pty.Window.Width,
			Height:     pty.Window.Height,
			time:       time.Now(),
			accOwner:   s.User(),
		}
		return login(m, tea.WithInput(s), tea.WithOutput(s), tea.WithAltScreen(), tea.WithMouseCellMotion())
	}
	return bm.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
}

type model struct {
	SSHSession ssh.Session
	time       time.Time
	Height     int
	Width      int
	accOwner   string // Account owner will be used for matching the characters created from this account.
}

type timeMsg time.Time

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timeMsg:
		m.time = time.Time(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "l", "ctrl+l":
			return mn.InitialModel(m.accOwner, m.SSHSession), nil // Go to the login page with passing account owner
		case "n", "ctrl+n":
			//mn.NewAccount()
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Welcome to gmud!\n\n"
	s += "Date> " + m.time.Format(time.RFC1123) + "\n\n"
	s += "Press 'l' to go in.\n"

	return fmt.Sprintln(s, m.Height, m.Width)
}
