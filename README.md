## Pokedex CLI

A fun project I've used to expand my knowledge on Go.

It's a repl based tool that utilises the incredible Pokemon API: https://pokeapi.co/ and allows you to explore the map, find and catch pokemon and inspect them in the pokedex.
### Features
Explore the Pokémon World:
Browse through paginated lists of Pokémon locations and discover new areas to explore.

Catch Pokémon:
Try your luck at catching Pokémon you encounter. Each catch attempt is randomized based on the Pokémon’s stats.

Pokédex Management:
View a list of all Pokémon you’ve caught and inspect their stats, types, height, and weight.

REPL Interface:
Interactive Read-Eval-Print Loop lets you enter commands and see results instantly.

Caching:
Uses an in-memory cache to reduce API calls and speed up repeated queries.

Commands
help — Show available commands and usage.
map — Display a paginated list of Pokémon locations.
mapb — Go back to the previous page of locations.
explore <location> — Explore a specific location to find Pokémon.
catch <pokemon> — Attempt to catch a Pokémon by name.
list-pokemon — List all Pokémon you’ve caught.
inspect <pokemon> — Show detailed info about a caught Pokémon.
exit — Exit the application.
