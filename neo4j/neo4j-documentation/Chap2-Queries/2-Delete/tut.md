

**1. delete a node**

```
MATCH (n:Person {name: 'Tom Hanks'}) 
DELETE n
```
result -- Deleted 1 node

It is also possible to delete the single node using the `NODETACH DELETE` clause. Using the NODETACH keyword
explicitly defines that relationships will not be detached and deleted from a node. T
``` 
MATCH (n:Person {name: 'Tom Hanks'}) 
NODETACH DELETE n
```

**2. delete a relation**

```
MATCH (n:Person {name: 'Laurence Fishburne'})-[r:ACTED_IN]->() 
DELETE r
```
result -- Deleted 1 relationship

**3. Delete a node with all its relationships**
```
MATCH (n:Person {name: 'Carrie-Anne Moss'}) 
DETACH DELETE n
```
result -- Deleted 1 node, deleted 1 relationship

**4. Delete all nodes and relationships**
```
MATCH (n) 
DETACH DELETE n
```
result -- Deleted 3 nodes, deleted 1 relationship
