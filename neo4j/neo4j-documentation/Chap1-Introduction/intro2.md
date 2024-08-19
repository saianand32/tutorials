

1. Neo4j consists of two editions: a commercial Enterprise Edition, and a Community Edition.

see --> intro2 image


2. Key Neo4j terminology --> intro3 image

3. `Cypher and Neo4j transactions` 
All Cypher queries run within transactions. Modifications done by updating queries are held in memory by
the transaction until it is committed, at which point the changes are persisted to disk and become visible to
other transactions. If an error occurs - either during query evaluation, such as division by zero, or during
commit, such as constraint violations - the transaction is automatically rolled back, and no changes are
persisted in the graph.
In short, an updating query always either fully succeeds or does not succeed at all.

4. ` Database Transactions & ACID properties` --> intro4 image

5. ` Cypher and Aura `
This section gives a brief overview of Aura, and how Cypher differs to users of Aura.
What is Aura?
Aura is Neo4j’s fully managed cloud service. It consists of AuraDB and AuraDS. AuraDB is a graph
database service for developers building intelligent applications, and AuraDS is a Graph Data Science
(GDS) service for data scientists building predictive models and analytics workflows.
AuraDB is available on the following tiers:
• AuraDB Free
• AuraDB Pro
• AuraDB Enterprise