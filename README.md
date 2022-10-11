# Connecting to a database (CouchDB)

`New` is used to create a database handle:

```go
client, err := kivik.New(context.TODO(), driver, dataSourceName)
```

Where `driver` specifies a database driver, and `dataSourceName` specifies database-specific connection information, such as a URL.

Note that open may or may not verify the existence of the database, permissions, etc. Depending on the driver, such checks may be deferred until a request is made.

## Authentication

Different drivers may support different authentication methods. The CouchDB driver (and the PouchDB driver for remote connections) supports auth credentials in the data source URL:

```go
client, err := kivik.New(context.GODO(), "couch", "http://admin:abc123@localhost/")
```

# User management

By default, users in CouchDB are stored as [user documents](http://docs.couchdb.org/en/2.1.0/intro/security.html?highlight=user%20documents#users-documents) in the `_users` database. This means user management is just a matter of manipulating these documents as usual.  For example, to create a new user:

```go
usersDB, _ := client.DB(context.TODO(), "_users") // Connect to the _users database
user := map[string]interface{}{
		"_id":      kivik.UserPrefix + "username",
		"type":     "user",
		"password": "abc123",
	}
rev, _ := usersDB.Put(context.TODO(), kivik.UserPrefix+"username", user)
```

# Storing a single document

Storing a document is done with `Put` or `Create` which correspond to [`PUT /{db}/{doc}`](http://docs.couchdb.org/en/2.0.0/api/database/common.html#put--db), [`POST /{db}`](http://docs.couchdb.org/en/2.0.0/api/database/common.html#post--db) respectively.  In most cases, you should use `Put`:

```go
type Animal struct {
    ID       string `json:"_id"`
    Rev      string `json:"_rev,omitempty"`
    Feet     int    `json:"feet"`
    Greeting string `json:"greeting"`
}

cow := Animal{ID: "cow", Feet: 4, Greeting: "moo"}
rev, err := db.Put(context.TODO(), "cow", cow)
if err != nil {
    panic(err)
}
cow.Rev = rev
```

# Updating a document

Updating a document is the same as storing one, except that the `_rev` parameter must match that stored on the server.

```go
cow.Rev = rev // Must be set
cow.Greeting = "Moo!"
newRev, err := db.Put(context.TODO(), "cow", cow)
if err != nil {
    panic(err)
}
cow.Rev = newRev
```

# Deleting a document

As with updating a document, deletion depends on the proper `_rev` parameter.

```go
newRev, err := db.Delete(context.TODO(), "cow", "2-9c65296036141e575d32ba9c034dd3ee")
if err != nil {
    panic(err)
}
fmt.Printf("The tombstone document has revision %s\n", newRev)
```

# Fetching a document

When fetching a document, the document will be unmarshaled from JSON into your structure by the `row.ScanDoc` method.

```go
row, err := db.Get(context.TODO(), "cow")
if err != nil {
    panic(err)
}
var cow Animal
if err = row.ScanDoc(&cow); err != nil {
    panic(err)
}
fmt.Printf("The cow says '%s'\n", cow.Greeting)
```

# Design documents

Design documents are treated identically to normal documents by Kivik, the only difference being the document ID.

- Create/update design documents [the same as normal documents](https://github.com/flimzy/kivik/wiki/Usage-Examples#storing-a-single-document)
- Delete a design document [the same as normal](https://github.com/flimzy/kivik/wiki/Usage-Examples#deleting-a-document) (although you may wish to [clean up views](https://godoc.org/github.com/flimzy/kivik#DB.ViewCleanup) after).

# Create or update a View

Store your document [normally](https://github.com/flimzy/kivik/wiki/Usage-Examples#storing-a-single-document), formatted with your views (or other functions).  Example:

    db.Put(context.TODO(), "_design/foo", map[string]interface{}{
        "_id": "_design/foo",
       "views": map[string]interface{}{
            "foo_view": map[string]interface{}{
                "view": "function(doc) { ... }",
            },
        },
    })

# Query a view

    rows, err := db.Query(context.TODO(), "_design/foo", "_view/bar", kivik.Options{
        "startkey": "foo",
        "endkey":   "foo" + kivik.EndKeySuffix,
    })
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        var doc interface{}
        if err := rows.ScanDoc(&doc); err != nil {
            panic(err)
        }
        /* do something with doc */
    }
    if rows.Err() != nil {
        panic(rows.Err())
    }

# Query a view with compound keys

    rows, err := db.Query(context.TODO(), "_design/foo", "_view/bar", kivik.Options{
        "startkey": []string{"foo", "bar"},
        "endkey":   []string{"foo", "bar" + kivik.EndKeySuffix},
    })
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        var doc interface{}
        if err := rows.ScanDoc(&doc); err != nil {
            panic(err)
        }
        /* do something with doc */
    }
    if rows.Err() != nil {
        panic(rows.Err())
    }

# Query a view literal JSON keys

    rows, err := db.Query(context.TODO(), "_design/foo", "_view/bar", kivik.Options{
        "startkey": json.RawMessage(`{"foo":true}`),
    })
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        var doc interface{}
        if err := rows.ScanDoc(&doc); err != nil {
            panic(err)
        }
        /* do something with doc */
    }
    if rows.Err() != nil {
        panic(rows.Err())
    }

# Map/Reduce

Use the [Query](https://godoc.org/github.com/flimzy/kivik#DB.Query) method with the appropriate options. See [Query a view](https://github.com/flimzy/kivik/wiki/Usage-Examples#query-a-view) for a more complete example of querying a view.

    rows, err := db.Query(context.TODO(), "_design/foo", "_view/bar", map[string]interface{}{"group": true})
