package gmud

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	//"github.com/charmbracelet/wish/bubbletea"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
	"github.com/muesli/termenv"
	"charactersroutes"
	"models"
)

const (
	host = "localhost"
	port = 2222
)

func main() {
	go charactersRoutes()

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			lm.Middleware(),
			func(h ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					mrj, _, _, _, _ := ssh.ParseAuthorizedKey(
						[]byte("ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMrr9hgSKnoddIDmzFyMnf5qb3QTsG40/9UyhexKiw6z mike@mikej.dev"),
					)
					switch {
					case ssh.KeysEqual(s.PublicKey(), mrj):
					 // TODO: Echo username, not ssh string
					default:
						wish.Println(s, "User not found!")
					}
					h(s)

				}
			},
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

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
		pty, _, active := s.Pty()
		if !active {
			wish.Fatalln(s, "no active terminal")
			return nil
	}
		m := model{
			term: pty.Term,
			width:    pty.Window.Width,
			height:    pty.Window.Height,
			time: time.Now(),
		}
		return login(m, tea.WithInput(s), tea.WithOutput(s)) //tea.WithAltScreen)
	}
	return bm.MiddlewareWithProgramHandler(teaHandler, termenv.ANSI256)
}

type model struct {
	term   string
	width  int
	height int
	time   time.Time
}

type timeMsg time.Time

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timeMsg:
		m.time = time.Time(msg)
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "l", "ctrl+l":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Welcome to gmud!\n"
	s += "Your terminal is: %s\n"
	s += "Your window size is x: %d y: %d\n\n"
	s += "The date is " + m.time.Format(time.RFC1123) + "\n\n"
	s += "Press l to login\n"
	s += "Press 'q' to quit\n"
	return fmt.Sprintf(s, m.term, m.width, m.height)
}

//TODO: Nothing happens after user logs in and their name is displayed