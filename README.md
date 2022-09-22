# cheevos

[![Test](https://github.com/haleyrc/cheevos/actions/workflows/go.yml/badge.svg)](https://github.com/haleyrc/cheevos/actions/workflows/go.yml)

Cheevos is an app to let people award custom achivements (Cheevos) to their
friends and colleages.

## Domain

### Organizations

Users are grouped together into Organizations. An organization represents a
boundary for cheevos. Every cheevo is owned by a single organization and can
only be awarded to members of that organization by the members of that same
organization.

Users can be members of multiple organizations and collect cheevos from all
of them, but the cheevos are tied to their status as members of that
organization. If they are removed from the organization in the future, the
cheevos they received from the organization are also removed.

## TODO

- [X] Refactor validation runner and compareError
- [X] Move membership to new service
- [X] Move awards to new service
- [X] Move password hash out of model
- [X] Implement an actual random string function
- [X] Combine mock repositories into one
- [X] Fix package stutter
- [X] Add invitation logger
- [X] Add award logger
- [X] Add membership logger
- [ ] Replace raw errors with error types
- [ ] Handlers
- [ ] Database
- [ ] Consistent variable naming after context merge
- [ ] Alphabetize methods
- [ ] Refactor creating test users and orgs in repo tests
