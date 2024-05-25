# Requirements:

Install templ as per [https://templ.guide/quick-start/installation]
> templ is useful to make templates for the html pages; Most pages will have similar components and this will help us easily copy paste html components

Install go air as per [https://github.com/cosmtrek/air]
> This makes developement easier; Instead of running multiple files after each change, we have hot reloading

Get the CLI of tailwind and place it in root; 
> I placed it in root because I was not sure where else to keep it


# Instructions
Run 
```bash
air -d
```
There should be a grid of rounded squares in [localhost:3000/about]
This will watch the folders for changes and host the page locally.
Sometimes, there is some compilation error, but saving the main.go again usually fixes it; i think that happens because of some timing issue.

# Explanations

Templ is really useful to make typsafe templates. Feels very intuitive to use it

While templ provides a way for generating css, it cannot generate other effects, like on-hover and so on. For this i decided to use tailwind instead, and with a nice lsp, this is not too bad.