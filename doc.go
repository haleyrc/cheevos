// Package cheevos provides a service layer for an achievement system for real
// people. Users have the ability to award custom Cheevos (which are similar to,
// but legally distinct from other achievements) to their friends and
// colleagues.
//
// Organizations
//
// Users are grouped together into Organizations. An organization represents a
// boundary for cheevos. Every cheevo is owned by a single organization and can
// only be awarded to members of that organization by the members of that same
// organization.
//
// Users can be members of multiple organizations and collect cheevos from all
// of them, but the cheevos are tied to their status as members of that
// organization. If they are removed from the organization in the future, the
// cheevos they received from the organization are also removed.
package cheevos
