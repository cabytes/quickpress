# Zine

Library for easily create blogging for your golang app or directly by using the CLI application on the fly. Access admin by navigating to `/admin`.

## Use with Chi

```
mux := chi.NewMux()

mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
})

app := zine.New(
    zine.DataPath("../../data"),
    zine.BaseHref("/blog"),
    zine.LoadTheme("../../themes/light/", light.Files),
    zine.AuthHook(func(username, password string) zine.User {
        if username == "zine" && password == "zine" {
            return &BlogUser{
                User:     user,
                Pass:     pass,
                UserName: "admin",
            }
        }
        return nil
    }),
)

mux.Handle("/blog*", app)

http.ListenAndServe(":8000", mux)
```

## From command line

```
zine --username=zine --password=zine --listen=:8000
```