.SILENT:

server:
	air

tailwind:
	tailwindcss -i themes/input.css -o themes/output.css --watch

deps:
	go install github.com/air-verse/air@latest && go install github.com/a-h/templ/cmd/templ@latest && go mod tidy

dev:
	make -j2 tailwind server
