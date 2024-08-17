
1. `What is Cypher` ??
    - Cypher is Neo4j’s declarative query language, allowing users to unlock the full potential of property graph databases.
    - Cypher is Neo4j’s declarative graph query language. It was created in 2011 by Neo4j engineers as an
      SQL-equivalent language for graph databases. Similar to SQL, Cypher lets users focus on what to retrieve
      from graph, rather than how to retrieve it. 
    - Cypher provides a visual way of matching patterns and relationships. It relies on the following ascii-art
      type of syntax: (nodes)-[:CONNECT_TO]→(otherNodes). 

2. `Cypher and SQL Key differences` - 

    a. `Cypher is schema-flexible` ---> While it is both possible and advised to enforce partial schemas using indexes and constraints, Cypher
                                      and Neo4j offers a greater degree of schema-flexibility than SQL and a relational database.

    b. `Query order` --> SQL queries begin with what a user wants to return, whereas Cypher queries end with the return clause. 

    c. `Cypher queries are more concise` ---> more easy and small to write especially for relational data-- below are same data fetch query in sql and cypher.
      
 ` SELECT actors.name FROM actors
   LEFT JOIN acted_in ON acted_in.actor_id = actors.id
   LEFT JOIN movies ON movies.id = acted_in.movie_id
   WHERE movies.title = "The Matrix" `

 `MATCH (actor:Actor)-[:ACTED_IN]->(movie:Movie {title: 'The Matrix'}) RETURN actor.name `


3. CYPHER and APOC(Awesome Procedures On Cypher)
    - Neo4j supports the APOC (Awesome Procedures on Cypher) Core library. The APOC Core library provides
      access to user-defined procedures and functions which extend the use of the Cypher query language into
      areas such as data integration, graph algorithms, and data conversion.



