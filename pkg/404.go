package whatever

import (
	"net/http"
)

func (Handler) Get404(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
    <title>Bummer</title>
    <link rel="icon" type="image/x-icon" href="/public/cat.favicon.ico">
</head>
<body>
<pre style="text-align: center;"><code>404: Not Found</code></pre>
</body>
</html>
`))
}
