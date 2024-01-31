ServerID ?= 1
Address ?= 127.0.0.1:13561

.PHONY: world
# gate-server
world:
	go run server/cmd/world_server/main.go

.PHONY: gate
# gate-server
gate:
	go run server/cmd/gate_server/main.go -s=$(ServerID) -ws=$(Address)

.PHONY: login
# login-server
login:
	go run server/cmd/login_server/main.go

.PHONY: game
# game-server
game:
	go run server/cmd/game_server/main.go -s=$(ServerID) -ws=$(Address)

.PHONY: client
# client for test
client:
	go run ./client/.
