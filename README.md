# neo4j


### Example Commands

```
// unwind

UNWIND [{name: "Steve", position: "Developer"},
  {name: "Mark", position: "Developer"},
  {name: "Jimmy", position: "Product Manager"}
] as user
CREATE (u:User {name: user.name, position: user.position})
```

```
// Create Relationship between Admin and all Groups with a Regex
MATCH (a:User),(b:Group)
WHERE a.name = "Admin" AND b.name =~ '.*'
CREATE (a)-[r:ADMIN_OF]->(b)
RETURN r
```

```
// Create Relationship between Steve and two groups
MATCH (a:User),(b:Alert)
WHERE a.name = "Steve" AND b.name IN ['group1', 'group2']
CREATE (a)-[r:VIEWER_OF]->(b)
RETURN 
```

