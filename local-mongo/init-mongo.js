// init user in mongo
db.createUser({
    user: "root",
    pwd: "password",
    roles: [{
        role: "readWrite",
        db: "todolist"
    }]
})

// createIndex 
// db.users.createIndex({ username: 1 }, { unique: true });

db.users.insertMany([
    {
        username: "admin",
        password: "admin",
        role: "admin"
    },
    {
        username: "user",
        password: "user",
        role: "user"
    }
])
