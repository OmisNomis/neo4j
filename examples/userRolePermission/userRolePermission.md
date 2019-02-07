# Simple User->Role->Page permission system

## Create Nodes

```
//users
CREATE (`Steve`:user {name:'Steve', mobile:'+4422222222'}),
 (`John`:user {name:'John', mobile:'+4422222222'}),
 (`Sarah`:user {name:'Sarah', mobile:'+4422222222'}),
 (`Liz`:user {name:'Liz', mobile:'+4422222222'}),
 (`Mark`:user {name:'Mark', mobile:'+4422222222'})
```

```
//pages
CREATE (`Page1`:page {name:'Page1'}),
 (`Page2`:page {name:'Page2'}),
 (`Page3`:page {name:'Page3'}),
 (`Page4`:page {name:'Page4'})
```

```
//roles 
CREATE (`Admin`:role {name:'Admin'}),
  (`Viewer`:role {name:'Viewer'}),
  (`Writer`:role {name:'Writer'}),
  (`Deleter`:role {name:'Deleter'})
```


## Relationships

```
//user-role relationships (Must be run at the same time as the above CREATES)
CREATE (`Steve`)-[:MEMBER_OF]->(`Admin`),
(`John`)-[:MEMBER_OF]->(`Admin`),
(`Sarah`)-[:MEMBER_OF]->(`Admin`),
(`Liz`)-[:MEMBER_OF]->(`Writer`),
(`Mark`)-[:MEMBER_OF]->(`Deleter`)
```

```
//user-role relationships (Can be run ad-hoc)
MATCH (a:user),(b:role)
WHERE a.name = "Mark" AND b.name = 'Viewer'
CREATE (a)-[r:MEMBER_OF]->(b)
RETURN r
```

```
// role-page relationships

// Admin can do everything
MATCH (a:role),(b:page)
WHERE a.name = "Admin" AND b.name =~ '.*'
CREATE (a)-[r:READ]->(b)
RETURN r

MATCH (a:role),(b:page)
WHERE a.name = "Admin" AND b.name =~ '.*'
CREATE (a)-[r:WRITE]->(b)
RETURN r

MATCH (a:role),(b:page)
WHERE a.name = "Admin" AND b.name =~ '.*'
CREATE (a)-[r:DELETE]->(b)
RETURN r

// Viewer can only READ
MATCH (a:role),(b:page)
WHERE a.name = "Viewer" AND b.name =~ '.*'
CREATE (a)-[r:READ]->(b)
RETURN r

MATCH (a:role),(b:page)
WHERE a.name = "Writer" AND b.name =~ '.*'
CREATE (a)-[r:WRITE]->(b)
RETURN r

MATCH (a:role),(b:page)
WHERE a.name = "Deleter" AND b.name =~ '.*'
CREATE (a)-[r:DELETE]->(b)
RETURN r
```


# Queries

```
// Get user-permission-page
MATCH paths=(u:user {name:'John'})-[:MEMBER_OF]->()-[r:READ|:WRITE]->(p2:page)
RETURN u.name AS User, TYPE(r), p2.name AS Page
```

```
// Delete relationships
MATCH ()-[r:CAN_VIEW|:DELETE]-() 
delete r
```
