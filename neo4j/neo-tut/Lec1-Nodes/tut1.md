1. Select all nodes
    $ Match ( n ) return n

2. Select node using label
    $ Match ( n: PLAYER ) return n

3. Select node using multiple labels
    $ Match ( n: PLAYER :GUARD ) return n

4. Select node properties and Alias them
    $ Match ( n: PLAYER ) return n.age as Age, n.height as Height

5. Select node using label and properties (2 ways)
    $ Match ( n: PLAYER {name: "LeBron James"}) return n
    $ Match ( n: PLAYER ) where n.name="LeBron James" return n

6. Not clause
    $ Match ( n: PLAYER ) where n.name <> "LeBron James" return n

7. Greater and Less than
    $ Match ( n: PLAYER ) where n.height >= 2.4  return n
    $ Match ( n: PLAYER ) where n.height <= 2.4  return n

8. Doing Math (BMI)
    $ Match ( n: PLAYER ) where (n.weight/(n.height * n.height)) > 25 return n

9. AND , OR
    $ Match ( n: PLAYER ) where n.height >= 2.4 AND n.age <30   return n
    $ Match ( n: PLAYER ) where n.height >= 2.4 OR n.age <30   return n

10. NOT (here the not will apply to the entire condition [ NOT of (this and that) ] )
    $ Match ( n: PLAYER ) where NOT n.height >= 2.4 AND n.age > 25  return n