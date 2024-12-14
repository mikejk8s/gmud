---
id: e6wea
title: Add db functions
file_version: 1.1.2
app_version: 1.8.5
---


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 README.md
```markdown
17     4. ~~DeleteUser currently 404s~~
18     5. CharacterRoutes GetUser returns the wrong list of characters, doesn't do its query
19     6. Require unique names sql queries Character.Name
20     7. Make a map of public keys
21     8. *AccountLogin function runs login page on the server*
22     9. Docker-compose broken, gmud connection refused to mysql/localhost:3306 - can access w/ sqlstudio fine
23     10. Alive reports false need to switch to reverse bool or dead
24     11. Level shouldnt be 0
25     12.
26     
27     
28     ## Api Paths
29     
30     Gin Stats http://127.0.0.1:8080/stats
31     http://127.0.0.1:8081/characters/9 {id}
32     
33     ``` go
34     	a := r.Group("/api")
35     	{
36     		a.POST("/token", controllers.GenerateToken)
37     		a.POST("/user/register", controllers.RegisterUser)
38     		r.GET("/characters", cr.GetCharacters)
39     		s := a.Group("/secured").Use(middlewares.Auth())
40     		{
41     			s.GET("/user", controllers.GetUser)
42     			s.POST("/token", controllers.GenerateToken)
43     			s.GET("/characters/:id", cr.GetCharacter)
44     			s.POST("/characters", cr.CreateCharacter)
45     			s.PUT("/characters/:id", cr.UpdateCharacters)
46     			s.DELETE("/characters/:id", cr.DeleteCharacter)
47     		}
48     ```
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 cmd/app/main.go
```go
20     	mn "github.com/mikejk8s/gmud/pkg/menus"
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 cmd/app/main.go
```go
48     	go tracing.JaegerTraceProvider()
49     
50     	// SSH server begin
51     	s, err := wish.NewServer(
52     		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
53     		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
54     		wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 cmd/app/main.go
```go
124    		return login(m, tea.WithInput(s), tea.WithOutput(s), tea.WithAltScreen())
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 cmd/app/main.go
```go
154    			mn.AccountLogin()
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 cmd/app/main.go
```go
168    	s += "Press 'l' to login\n"
```

<br/>

This code defines a command line interface in Go using the Bubble Tea library. It allows the user to select an item from a list using the up and down arrow keys and toggle its selection using the spacebar or enter key. The selected items are tracked using a map and the program can be exited using the ctrl+c or q keys.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 extra/charselect.go
```go
1      package changeme
2      
3      import (
4      	"fmt"
5      	tea "github.com/charmbracelet/bubbletea"
6      	"os"
7      )
8      
9      type model struct {
10     	choices  []string         // items on the to-do list
11     	cursor   int              // which to-do list item our cursor is pointing at
12     	selected map[int]struct{} // which to-do items are selected
13     }
14     
15     func initialModel() model {
16     	return model{
17     		// Our shopping list is a grocery list
18     		choices: []string{"Gandalf", "Fender", "Ghibli"},
19     
20     		// A map which indicates which choices are selected. We're using
21     		// the  map like a mathematical set. The keys refer to the indexes
22     		// of the `choices` slice, above.
23     		selected: make(map[int]struct{}),
24     	}
25     }
26     
27     func (m model) Init() tea.Cmd {
28     	// Just return `nil`, which means "no I/O right now, please."
29     	return nil
30     }
31     
32     func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
33     	switch msg := msg.(type) {
34     
35     	// Is it a key press?
36     	case tea.KeyMsg:
37     
38     		// Cool, what was the actual key pressed?
39     		switch msg.String() {
40     
41     		// These keys should exit the program.
42     		case "ctrl+c", "q":
43     			return m, tea.Quit
44     
45     		// The "up" and "k" keys move the cursor up
46     		case "up", "k":
47     			if m.cursor > 0 {
48     				m.cursor--
49     			}
50     
51     		// The "down" and "j" keys move the cursor down
52     		case "down", "j":
53     			if m.cursor < len(m.choices)-1 {
54     				m.cursor++
55     			}
56     
57     		// The "enter" key and the spacebar (a literal space) toggle
58     		// the selected state for the item that the cursor is pointing at.
59     		case "enter", " ":
60     			_, ok := m.selected[m.cursor]
61     			if ok {
62     				delete(m.selected, m.cursor)
63     			} else {
64     				m.selected[m.cursor] = struct{}{}
65     			}
66     		}
67     	}
68     
69     	// Return the updated model to the Bubble Tea runtime for processing.
70     	// Note that we're not returning a command.
71     	return m, nil
72     }
73     
74     func (m model) View() string {
75     	// The header
76     	s := "Which character would you like to login as?\n\n"
77     
78     	// Iterate over our choices
79     	for i, choice := range m.choices {
80     
81     		// Is the cursor pointing at this choice?
82     		cursor := " " // no cursor
83     		if m.cursor == i {
84     			cursor = ">" // cursor!
85     		}
86     
87     		// Is this choice selected?
88     		checked := " " // not selected
89     		if _, ok := m.selected[i]; ok {
90     			checked = "x" // selected!
91     		}
92     
93     		// Render the row
94     		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
95     	}
96     
97     	// The footer
98     	s += "\nPress q to quit.\n"
99     
100    	// Send the UI for rendering
101    	return s
102    }
103    
104    func main() {
105    	p := tea.NewProgram(initialModel())
106    	if err := p.Start(); err != nil {
107    		fmt.Printf("Alas, there's been an error: %v", err)
108    		os.Exit(1)
109    	}
110    }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 extra/classselect.go
```go
1      // package menus
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 extra/classselect.go
```go
13     // }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/charactersroutes/charactersroutes.go
```go
8      	db "github.com/mikejk8s/gmud/pkg/mysql"
9      	)
10     
11      var Characters = []m.Character{
12     // 	{ID: "1", Name: "John Doe", Class: "Warrior", Race: "Human", Level: 1},
13     }
14     
15     func GetCharacters(c *gin.Context) {
16     	name := c.Param("name")
17     	for _, Character := range Characters {
18     		if Character.Name == name {
19     	c.JSON(http.StatusOK, Characters)
20     	db.GetCharacters(Character.Name)
21     			}
22     		}
23     		c.JSON(http.StatusNotFound, gin.H{"error": "Character Names not found"})
24     }
25     
26     
27     func GetCharacter(c *gin.Context) {
28     	id := c.Param("id")
29     	for _, Character := range Characters {
30     		if Character.ID == id {
31     			c.JSON(http.StatusOK, Character)
32     			//return
33     			db.GetCharacters(id)
34     		}
35     	}
36     	c.JSON(http.StatusNotFound, gin.H{"error": "Character ID not found"})
37     }
38     
39     func CreateCharacter(c *gin.Context) {
40     	var Character m.Character
41     	c.BindJSON(&Character)
42     	Characters = append(Characters, Character)
43     	c.JSON(http.StatusCreated, Character)
44     	db.AddCharacter(Character)
45     }
46     
47     func UpdateCharacters(c *gin.Context) {
48     	id := c.Param("id")
49     	var Character m.Character
50     	c.BindJSON(&Character)
51     	for index, item := range Characters {
52     		if item.ID == id {
53     			Characters[index] = Character
54     			c.JSON(http.StatusOK, Character)
55     			//return
56     		}
57     	}
58     	c.JSON(http.StatusNotFound, errors.New("Character not found"))
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/charactersroutes/charactersroutes.go
```go
44     	db.AddCharacter(Character)
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/charactersroutes/charactersroutes.go
```go
55     			//return
56     		}
57     	}
58     	c.JSON(http.StatusNotFound, errors.New("Character not found"))
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/menus/accountmenu.go
```go
1      package menus
2      
3      import (
4      	"fmt"
5      	"os"
6      
7      	tea "github.com/charmbracelet/bubbletea"
8      	//"github.com/charmbracelet/wish"
9      )
10     
11     type model struct {
12     	choices  []string         // items on the list
13     	cursor   int              // item our cursor is pointing at
14     	selected map[int]struct{} // whats selected
15     }
16     
17     func initialModel() model {
18     	return model{
19     		choices: []string{"Login", "Create Account", "Test"},
20     
21     		// A map which indicates which choices are selected. We're using
22     		// the  map like a mathematical set. The keys refer to the indexes
23     		// of the `choices` slice, above.
24     		selected: make(map[int]struct{}),
25     	}
26     }
27     
28     func (m model) Init() tea.Cmd {
29     	// Just return `nil`, which means "no I/O right now, please."
30     	return nil
31     }
32     
33     func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
34     	switch msg := msg.(type) {
35     
36     	// Is it a key press?
37     	case tea.KeyMsg:
38     
39     		// Cool, what was the actual key pressed?
40     		switch msg.String() {
41     
42     		// These keys should exit the program.
43     		case "ctrl+c", "q":
44     			return m, tea.Quit
45     
46     		// The "up" and "k" keys move the cursor up
47     		case "up", "k":
48     			if m.cursor > 0 {
49     				m.cursor--
50     			}
51     
52     		// The "down" and "j" keys move the cursor down
53     		case "down", "j":
54     			if m.cursor < len(m.choices)-1 {
55     				m.cursor++
56     			}
57     
58     		// The "enter" key and the spacebar (a literal space) toggle
59     		// the selected state for the item that the cursor is pointing at.
60     		case "enter", " ":
61     			_, ok := m.selected[m.cursor]
62     			if ok {
63     				delete(m.selected, m.cursor)
64     			} else {
65     				m.selected[m.cursor] = struct{}{}
66     			}
67     		}
68     	}
69     
70     	// Return the updated model to the Bubble Tea runtime for processing.
71     	// Note that we're not returning a command.
72     	return m, nil
73     }
74     
75     func (m model) View() string {
76     	// The header
77     	s := "Login or create account?\n\n"
78     
79     	// Iterate over our choices
80     	for i, choice := range m.choices {
81     
82     		// Is the cursor pointing at this choice?
83     		cursor := " " // no cursor
84     		if m.cursor == i {
85     			cursor = ">" // cursor!
86     		}
87     
88     		// Is this choice selected?
89     		checked := " " // not selected
90     		if _, ok := m.selected[i]; ok {
91     			checked = "x" // selected! //TODO: Make this return a different page not just se
92     		}
93     
94     		// Render the row
95     		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
96     	}
97     
98     	// The footer
99     	s += "\nPress q to quit.\n"
100    
101    	// Send the UI for rendering
102    	return s
103    }
104    
105    // func AccountLogin() {
106    // 	p := tea.NewProgram(initialModel())
107    // 	if err := p.Start(); err != nil {
108    // 		fmt.Printf("Alas, there's been an error: %v", err)
109    // 		os.Exit(1)
110    // 	}
111    // 	}
112    
113    func AccountLogin() error {
114    	p := tea.NewProgram(initialModel())
115    	if err := p.Start(); err != nil {
116    		fmt.Printf("Alas, there's been an error: %v", err)
117    		os.Exit(1)
118    	}
119    	return nil
120    }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/models/character.go
```go
3      import (
4      	"time"
5      )
6      
7      type Character struct {
8      	Name string	`json:"name"`
9      	ID string	`json:"id"`
10     	Class string	`json:"class"`
11     	Race string	`json:"race"`
12     	Level int	`json:"level"`
13     	CreatedAt time.Time	`json:"created_at"`
14     	Alive bool	`json:"alive"`
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/charactersdb.go
```go
14     // func GetCharacters(code string) []m.Character {
15     // 	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
16     // 	char:= &m.Character{}
17     // 	if err != nil {
18     // 		fmt.Println("Error", err.Error())
19     // 		return nil
20     // 		{
21     // 	defer db.Close()
22     // 	results, err := db.Query("SELECT * FROM characters")
23     // 	if err != nil {
24     // 		fmt.Println("Err", err.Error())
25     // 		return nil
26     // 	}
27     // 	defer results.Close()
28     // 	for result.Next()
29     
30     func GetCharacter() []m.Character {
31     	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?parseTime=true")
32     	if err != nil {
33     		fmt.Println("Error", err.Error())
34     		return nil
35     }
36     
37     defer fgtrace.Config{Dst: fgtrace.File("charactersdb-fgtrace.json")}.Trace().Stop()
38     
39     defer db.Close()
40     results, err := db.Query("SELECT * FROM characters")
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/charactersdb.go
```go
49     			err = results.Scan(&character.Name, &character.ID, &character.Class, &character.Race, &character.Level, &character.CreatedAt, &character.Alive)
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/charactersdb.go
```go
58     				func GetCharacters(code string) *m.Character {
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/charactersdb.go
```go
68     					results, err := db.Query("SELECT * FROM characters WHERE id = ?", code)
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/charactersdb.go
```go
93     						"INSERT INTO characters (name,id,class,level,race) VALUES (?,?,?,?,?)",
94     						Character.Name, Character.ID, Character.Class, Character.Level, Character.Race)
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/charactersdb.go
```go
102    				func DeleteCharacter(Character m.Character) {
103    					db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname)
104    					if err!= nil {
105    						panic(err.Error())
106    						}
107    						// defer the close till after this function has finished
108    						// executing
109    						defer db.Close()
110    						delete, err := db.Query(
111    							"DELETE FROM characters WHERE id = ?", Character.ID)
112    							// if there is an error deleting, handle it
113    							if err!= nil {
114    								panic(err.Error())
115    								}
116    								defer delete.Close()
117    								}
118    
119    
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/mysql/createdb.go
```go
75     		race VARCHAR(15) NOT NULL DEFAULT 'HUMAN',
76     		level INT(3) NOT NULL DEFAULT '1',
```

<br/>

This file was generated by Swimm. [Click here to view it in the app](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZ211ZCUzQSUzQW1pa2Vqazhz/docs/e6wea).
