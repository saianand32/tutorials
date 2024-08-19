

1. find nodes and type alias
`MATCH (keanu:Person {name:'Keanu Reeves'}) RETURN keanu.name AS name, keanu.born AS born`

2. find nodes with limit & skip
`MATCH (people:Person) RETURN people SKIP 1 LIMIT 5`

3. find connected nodes
`MATCH (m:Movie {title: 'The Matrix'})<-[d:DIRECTED]-(p:Person) RETURN p.name as director`


4. Finding paths ***VVVIMP***
-- find path between two nodes Person and collegue 2 hops away persons name is Tom Hanks

` MATCH p = (n:Person {name: 'Tom Hanks'})--{2}(colleague:Person) RETURN p LIMIT 5`

-- find path between two nodes Person and collegue 1 to 6 hops away persons name is Tom Hanks

` MATCH p = (n:Person {name: 'Tom Hanks'})--{1, 6}(colleague:Person) RETURN p LIMIT 5`


5. Shrtest path
The SHORTEST keyword can be used to find a variation of the shortest paths between two nodes. In this
example, ALL SHORTEST paths between the two nodes Keanu Reeves and Tom Cruise are found. The
count() function calculates the number of these shortest paths while the length() function calculates the
length of each path in terms of traversed relationships.
Query

//returns all shortest paths as pathength might be same for the shortest ones
` MATCH p = ALL SHORTEST (:Person {name:"Keanu Reeves"})--+(:Person {name:"Tom Cruise"}) RETURN count(p) AS pathCount, length(p) AS pathLength`

//just 1 shortest path
` MATCH p = SHORTEST 1 (:Person {name:"Keanu Reeves"})--+(:Person {name:"Tom Cruise"}) RETURN count(p) AS pathCount, length(p) AS pathLength`