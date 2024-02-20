# Clearance

Open source siem built from scratch.

# Stack

Stack is comprised of postgres, go and htmx, plus a bit of css.
It will support every feature 


# Why

SIEMS are slow and terrible to use. We can fix that.

# Requirements

Postgres DB
GO v1.22 (previous versions might be ok too, idk, probably)

# Installation

```
git clone https://github.com/Woshiwuja/clearance
cd clearance
go mod tidy
cd server
go build main.go
./main
```

# Current features/TODO

- [ ] Log Parse
- [x] Add Devices
- [x] Add new rule (unused but you can add them lol)
- [x] Dashboarding (earlyyyyyyyyy)
