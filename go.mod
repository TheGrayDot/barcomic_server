module thegraydot.io/barcomic_server

go 1.18

require internal/server v0.2.0

require (
	github.com/mdp/qrterminal v1.0.1 // indirect
	github.com/micmonay/keybd_event v1.1.1 // indirect
	rsc.io/qr v0.2.0 // indirect
)

replace internal/server => ./internal/server
