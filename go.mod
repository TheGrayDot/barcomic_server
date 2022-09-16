module thegraydot.io/barcomic_server

go 1.18

require internal/server v1.0.0

require github.com/micmonay/keybd_event v1.1.1 // indirect

replace internal/server => ./internal/server
