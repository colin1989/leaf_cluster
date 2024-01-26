.PHONY: gate
# gate-server
gate:
	go run server/cmd/gate_server/main.go

.PHONY: login
# login-server
login:
	go run server/cmd/login_server/main.go

.PHONY: game
# game-server
game:
	go run server/cmd/game_server/main.go

.PHONY: client
# client for test
client:
	go run ./client/.
