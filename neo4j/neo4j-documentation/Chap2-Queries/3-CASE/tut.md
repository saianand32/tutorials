
# Conditional expressions (CASE)

Generic conditional expressions can be expressed in Cypher using the CASE construct. Two variants of CASE
exist within Cypher: the simple form, to compare a single expression against multiple values, and the
generic form, to express multiple conditional statements.

**CASE can only be used as part of RETURN or WITH if you want to use the result in a
subsequent clause.**


**1. Simple CASE**
The simple CASE form is used to compare a single expression against multiple values, and is analogous to
the switch construct of programming languages. The expressions are evaluated by the WHEN operator until
a match is found. If no match is found, the expression in the ELSE operator is returned. If there is no ELSE
case and no match is found, null will be returned.

 ``` MATCH (n:Person)
  RETURN
  CASE n.eyes
  WHEN 'blue' THEN 1
  WHEN 'brown', 'hazel' THEN 2
  ELSE 3
  END AS result, n.eyes
```


**2. The extended simple CASE** 
form allows the comparison operator to be specified explicitly.
The simple CASE
uses an implied equals (=) comparator.
The supported comparators are:
• Regular Comparison Operators: =, <>, <, >, <=, >=
• IS NULL Operator: IS [NOT] NULL
• Type Predicate Expression: IS [NOT] TYPED <TYPE> (Note that the form IS [NOT] :: <TYPE> is not
20
accepted)
• Normalization Predicate Expression: IS [NOT] NORMALIZED
• String Comparison Operators: STARTS WITH, ENDS WITH, =~ (regex matching)

```
MATCH (n:Person)
RETURN n.name,
CASE n.age
  WHEN IS NULL, IS NOT TYPED INTEGER | FLOAT THEN "Unknown"
  WHEN = 0, = 1, = 2 THEN "Baby"
  WHEN <= 13 THEN "Child"
  WHEN < 20 THEN "Teenager"
  WHEN < 30 THEN "Young Adult"
  WHEN > 1000 THEN "Immortal"
  ELSE "Adult"
END AS result
```

**3. Generic CASE**
The generic CASE expression supports multiple conditional statements, and is analogous to the if-elseif21
else construct of programming languages. Each row is evaluated in order until a true value is found. If no
match is found, the expression in the ELSE operator is returned. If there is no ELSE case and no match is
found, null will be returned.

```
MATCH (n:Person)
RETURN
CASE
  WHEN n.eyes = 'blue' THEN 1
  WHEN n.age < 40 THEN 2
  ELSE 3
END AS result, n.eyes, n.age
```

**4. CASE with null values**
When working with null values, you may be forced to use the generic CASE form. The two examples below
use the age property of the Daniel node (which has a null value for that property) to clarify the difference

```
MATCH (n:Person)
RETURN n.name,
CASE
  WHEN n.age IS NULL THEN -1
  ELSE n.age - 10
END AS age_10_years_ago
```

**5. CASE expressions and succeeding clauses**
The results of a CASE expression can be used to set properties on a node or relationship.

```
MATCH (n:Person)
WITH n,
CASE n.eyes
  WHEN 'blue' THEN 1
  WHEN 'brown' THEN 2
  ELSE 3
END AS colorCode
SET n.colorCode = colorCode
RETURN n.name, n.colorCode
```