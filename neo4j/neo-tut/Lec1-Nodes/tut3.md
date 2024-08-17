

1. using limit & skip
` MATCH (p:PLAYER) RETURN p skip 3 limit 10 `

2. Using labels to get all labels 
 ` MATCH(p:PLAYER)- ->(x) return labels(p)` 