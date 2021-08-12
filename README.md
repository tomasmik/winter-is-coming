# Winter is coming

[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://spacemacs.org)


[Winter is coming](https://github.com/mysteriumnetwork/winter-is-coming)

## Run 

Make sure your go install is [correctly configured](https://golang.org/doc/install#testing) then and run `make run` or build the binary yourself with `go build cmd/main.go` and run it.

To run the tests you can run `make test`.

## Interaction

Interacting with the server can be done with any number of tools, I chose `netcat`.
Commands the server understands:

```
# Join a server with a player name 
JOINSERVER {player}
```

```
# Join/Create game (if a doesn't exist, it'll get created)
JOINGAME {gameName}
```

```
# Shoot the zombie 
SHOOT {shoot}
```

I deviated a little bit from the given example, as it said itself that the given communication
is just an example. This made more sense to me.
