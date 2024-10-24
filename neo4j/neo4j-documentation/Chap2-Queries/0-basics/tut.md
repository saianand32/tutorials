# Core concepts

--- Fundamentally, a Neo4j graph database consists of three core entities: `nodes, relationships, and paths`.
Cypher queries are constructed to either match or create these entities in a graph. Having a basic
understanding of what nodes, relationships, and paths are in a graph database is therefore crucial in order
to construct Cypher queries.

1. `Nodes`
The data entities in a Neo4j graph database are called nodes. Nodes are referred to in Cypher using
parentheses `()`.

eg -- 
```
 MATCH (n:Person {name:'Anna'}) 
 RETURN n.born AS birthYear
 ```

In the above example, the node includes the following:
• A Person label. Labels are like tags, and are used to query the database for specific nodes. A node may
have multiple labels, for example Person and Actor.
• A name property set to Anna. Properties are defined within curly braces, {}, and are used to provide
nodes with specific information, which can also be queried for and further improve the ability to
pinpoint data.
• A variable, n. Variables allow referencing specified nodes in subsequent clauses.

2. `Relationships`
Nodes in a graph can be connected with relationships. A relationship must have a start node, an end node,
and exactly one type. Relationships are represented in Cypher with arrows (e.g. -->) indicating the
direction of a relationship.
eg -- 
```
MATCH (:Person {name: 'Anna'})-[r:KNOWS WHERE r.since < 2020]->(friend:Person) 
RETURN count(r) As numberOfFriends
```


3. `Paths`
Paths in a graph consist of connected nodes and relationships. Exploring these paths sits at the very core
of Cypher.

 ```
  MATCH (n:Person {name: 'Anna'})-[:KNOWS]-{1,5}(friend:Person WHERE n.born < friend.born) 
  RETURN DISTINCT friend.name AS olderConnections 
  ```
This example uses a quantified relationship to find all paths up to 5 hops away, traversing only
relationships of type KNOWS from the start node Anna to other older Person nodes (as defined by the
WHERE clause). The DISTINCT operator is used to ensure that the RETURN clause only returns unique
nodes.

```
MATCH p = SHORTEST 1 (:Person {name: 'Anna'})-[:KNOWS]-+(:Person {nationality: 'Canadian'})
```

Paths can also be assigned variables. For example, the above query binds a whole path pattern, which
matches the SHORTEST path from Anna to another Person node in the graph with a nationality property set
to Canadian. In this case, the RETURN clause returns the full path between the two nodes.

