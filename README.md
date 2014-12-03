# md2html

## building

```bash
go build .
```

### usage

Internally, *md2html* uses Go's standard html/template library. It provides a new function to the template: getMarkdown.

This function reads a markdown file from the location you specify in the template, converts it into HTML, and then inserts it into the HTML.

The CLI command that this code produces reads the template from STDIN and outputs to STDOUT.

#### Example

README.md
```markdown
# Test
This is some copy.
```

index.html.template
```html
<!doctype html>
<html>
<head>
	<meta charset="utf-8"/>
	<title>My README... but as a HTML file!</title>
</head>

<body>
	<div id="master">
		{{ getMarkdown "README.md" }}
	</div><!-- /#master -->
</body>
</html>
```

run command:
```bash
./md2html < index.html.template > index.html
```

outputs to index.html
```html
<!doctype html>
<html>
<head>
	<meta charset="utf-8"/>
	<title>My README... but as a HTML file!</title>
</head>

<body>
	<div id="master">
		<h1>Test</h1>
		<p>This is some copy.</p>
	</div><!-- /#master -->
</body>
</html>
```
