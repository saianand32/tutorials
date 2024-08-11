
1. Skip & Limit
    $ Match ( n: PLAYER ) where n.height >= 2.4  return n  SKIP 3  LIMIT 5

2. Sorting Data (Order by)
    $ Match ( n: PLAYER )  return n ORDER BY n.age ASC

3. Multiple nodes return (forms cartesian prod)
    $Match (player: PLAYER), (coach: COACH) where player.age > 12 return player, coach
