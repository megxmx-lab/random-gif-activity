package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	gifs := []string{
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatFeedingFranzy.gif",
		"https://s3.amazonaws.com/files.656.mba/mgt656/fall-2026/random-gifs/adorbs/CatPawPsychadelia.gif",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		count := 1

		countText := r.URL.Query().Get("count")
		if parsedCount, err := strconv.Atoi(countText); err == nil &&
			parsedCount >= 1 && parsedCount <= 3 {
			count = parsedCount
		}

		log.Printf("requested GIF count: %d", count)
		images := ""

		for i := 0; i < count; i++ {
			gifURL := gifs[rand.Intn(len(gifs))]
			images += fmt.Sprintf(
				`<img src="%s" alt="A cute animal">`,
				gifURL,
			)
		}

		page := fmt.Sprintf(`
<!doctype html>
<html>
<head>
    <title>Random Animal GIFs</title>
</head>
<body style="background-color: lightblue;">
    <h1>Hello, world! It is fall now.</h1>
	<form method="get" action="/">
    <label for="count">How many GIFs?</label>
    <select id="count" name="count">
        <option value="1">One</option>
        <option value="2">Two</option>
        <option value="3">Three</option>
    </select>
    <button type="submit">Show GIFs</button>
</form>
	%s
</body>
</html>
`, images)

		if _, err := w.Write([]byte(page)); err != nil {
			log.Printf("write response: %v", err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Open http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
