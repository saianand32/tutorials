1. querying nodes using a relationship
<!-- this returns all players playing for LA Lakers team -->
$ `Match(player:PLAYER) -[:PLAYS_FOR]-> (team:TEAM) where team.name = "LA Lakers"
   RETURN player, team`


2. Note like nodes have properties relations also can have properties
% here we can name the relation as contract and query as properties
$ `Match(player:PLAYER) -[contract:PLAYS_FOR]-> (team:TEAM) where team.name = "LA Lakers" AND contract.salary > 30000000
   return player,contract, team`


3. Write a query to print all players who are teammates with LeBron James and have salary > 30000000 

2 ways - 
a. `Match(:TEAM)<-[contract:PLAYS_FOR]-(player:PLAYER)<-[teamMates:TEAMMATES]-(player2:PLAYER {name: "LeBron James"}) 
   where contract.salary > 30000000
   return player`

b. By using multiple match statements
    `Match(teamMate:PLAYER)<-[:TEAMMATES]-(player:PLAYER{name:"LeBron James"})
    Match(teamMate)-[contract:PLAYS_FOR]->(team:TEAM)
    where contract.salary>30000000
    return teamMate`

