package main

const tpl = `
	<!DOCTYPE html>
	<html>
	<head><title>{{.Title}}</title></head>
	<body>
		<h2>{{.Header}}</h2>
		<ul>
			{{range $k, $v := .RecordsCache}}
				<li><strong>{{$k}}:</strong> {{$v}}</li>
			{{end}}
		</ul>
		{{if .RecrodsDB }}
		<h2>Records from db</h2>
		<ul>
			{{range $k, $v := .RecordsDB}}
				<li><strong>{{$k}}:</strong> {{$v}}</li>
			{{end}}
		</ul>
		{{ end}}
	</body>
	</html>
	`
