

Aggregation methods -- COUNT(), AVG(), SUM(), MAX(), MIN(), COLLECT()

1. return players and number of games played
    `Match(p:PLAYER)-[gamesPlayed:PLAYED_AGAINST]->(:TEAM) return p.name, count(gamesPlayed)`

2. return players and average of games points
    `Match(p:PLAYER)-[gamesPlayed:PLAYED_AGAINST]->(:TEAM) return p.name, avg(gamesPlayed.points)`

3. Return all players ages as an array
    ` MATCH(p: PLAYER) return COLLECT(p.age) ` -->> output - [33, 23, 44]