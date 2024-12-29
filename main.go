package main

import "fmt"
import "strconv"
import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            count, _ := strconv.Atoi(r.FormValue("counter"))
            count++
           fmt.Fprintf(w, `
            <body>
                <form action="/" method="POST">
                    <label>Counter</label>
                    <input name="counter" value="%d" readonly>
                    <button type="submit">Add</button>
                </form>
				<a href="/">Reset</a>
            </body>
        `, count)
		return
        }

        fmt.Fprintf(w, `
            <body>
                <form action="/" method="POST">
                    <label>Counter</label>
                    <input name="counter" value="1" readonly>
                    <button type="submit">Add</button>
                </form>
            </body>
            <a href="/">Reset</a>
        `)
    	})

	http.ListenAndServe(":80", nil)
}

/* Step 2
package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("q") != "" {
			fmt.Fprintf(w, `
				<body>
					<h1>Hello, %s</h1>
				</body>
			`, r.FormValue("q"))
			return
		}
		fmt.Fprintf(w, `
			<body>
				<form action="/" method="GET">
					<label>Enter your name</label>
					<input name="q">
					<button type="submit">Submit</button>
				</form>
			</body>
		`)
	})

	http.ListenAndServe(":80", nil)
}

/* Step 1
package main

import "net/http"
import "fmt"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "<h1>Hello, World!</h1>")
    })

    http.ListenAndServe(":80", nil)
}

/* Step 0
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
*/