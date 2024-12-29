# Getting Started with Go and the Web: Hello, World!

![Cover](https://res.cloudinary.com/craigsims808/image/upload/v1735440495/freelance/sturdy-train/cover_tdjtnd.png)

Golang is growing in popularity as a server-side language. Developers use it as an alternative to more common technologies like PHP, Node.js, Ruby, Python, etc. One of the appeals of Go is its low resource usage. Developers on a budget can build powerful apps and host them cheaply on low-end infrastructure if they use Go.

This guide aims to gently introduce Go and see how it can be used to build applications for the web. This tutorial is aimed at beginners; however, experienced developers can also browse through it. You never know—you might learn a thing or two. Enough with the chit-chat, let's get started.

## Prerequisites

To follow along with the guide, you need to have Go installed on your machine. Download and install Go by following the instructions: [Download and Install Go](https://go.dev/doc/install).

Some knowledge of HTML is required for you to understand the concepts being illustrated. HTML is easy to catch on if you haven't written it before. Follow this guide if you haven't got a clue about HTML [Learn HTML](https://developer.mozilla.org/en-US/docs/Web/HTML)

Some knowledge of HTTP is required for you to understand the guide better. HTTP is the transport protocol for the web. Follow this guide to understand HTTP [Learn HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP)

## Hello World, Go

To get started with Go, open up your work folder and create a new file named `main.go`. Add the following code to `main.go`
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Save the file and open up a new terminal session in your work directory. Run the following command:
```sh
go run main.go
```

You should see the following output:
```sh
Hello, World!
```

Congratulations! You have created your first program using Go. Pause and absorb the nice feeling. :) Now back to work. 

That was just a warm up to get things started.

## Hello World, Go on the Web

It's all well and good. We were able to print out "Hello, World!" on the terminal using Go, but that's not the aim of this guide. The real aim is to print out "Hello, World!" on a webpage using Go. So how do we do it?

First off, comment out the code you just wrote in `main.go`. At the top of your comment block, give it a ubiquitous name like "Step 1". 

Above the comment block, add the following code:

```go
package main

import "net/http"
import "fmt"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "<h1>Hello, World!</h1>")
    })

    http.ListenAndServe(":80", nil)
}
```

Save the file and run it. Visit `localhost` in your browser and you should see the text "Hello, World!" on a webpage.

![Hello World web page](https://res.cloudinary.com/craigsims808/image/upload/v1735436523/freelance/sturdy-train/hello-world-web-page_vrhx5b.png)

## Dynamic Hello

Awesome! We just managed to create a web server using Go and print the text "Hello, World!". The next step is to make the text change based on user input. Why? So that we understand the purpose of server side languages and how they can make our web pages dynamic.

Comment out the code you wrote in the previous step and at the top of the comment block give it a useful name like "Step 2".

Above the comment block add the following code:

```go
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
```
Stop your server (CTRL + C). Save the file and run it. Visit `localhost` in your browser and you should see a form input with the label, "Enter your name" and a "Submit" button. 

![Dynamic Web Page Go Form Input](https://res.cloudinary.com/craigsims808/image/upload/v1735436886/freelance/sturdy-train/dynamic-web-page-form-input_afljou.png)

Enter your name and click on "Submit". You should see the text "Hello, Jared", if your name is "Jared". Otherwise you will see the name you entered in the form.

![Dynamic Web Page Go Form Result](https://res.cloudinary.com/craigsims808/image/upload/v1735436887/freelance/sturdy-train/dynamic-web-page-form-result_rkdwvz.png)

## Counter

Next, let's try something a different. How about a counter? 

Comment out the code you wrote in the previous step and at the top of the comment block give it a useful name like "Step 3".

Above the comment block add the following code:

```go
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
```

Save the file and run it. Visit `localhost` in your browser and you should see a form input with the number "1" inside, the label, "Counter", and an "Add" button.

![Dynamic Web Page Counter](https://res.cloudinary.com/craigsims808/image/upload/v1735437587/freelance/sturdy-train/dynamic-web-page-counter_uufrna.png)

Clicking the button should increase the value of the counter.

## Conclusion

Congratulations! You've successfully built your first web application with Go, starting with a simple "Hello, World!" and progressing to a dynamic web page with user interaction. You've learned how to set up a basic web server, handle user input, and create a counter with Go. This is just the beginning—Go’s simplicity and efficiency make it a great choice for building scalable web applications.

As you continue exploring Go, you'll find many more advanced features and libraries that can help you build powerful and efficient applications. Keep experimenting and building, and soon you'll be mastering Go and creating complex web applications in no time. Happy coding!

Be on the look out for more tutorials. Ciao!

## Resources

- [GitHub Repo](https://github.com/Marktawa/sturdy-train)
- [Learn HTML](https://developer.mozilla.org/en-US/docs/Web/HTML)
- [Learn HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [Learn Go](https://go.dev/doc/)