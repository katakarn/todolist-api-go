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

db.todos.insertMany([
    // {
    //     username: "admin",
    //     password: "admin",
    //     role: "admin"
    // },
    // {
    //     username: "user",
    //     password: "user",
    //     role: "user"
    // },
    {
        title: "test1",
        description: "des1",
        isDone: false,
    },
    {
        title: "test2",
        description: "des2",
        isDone: false,
    },
    {
        title: "test3",
        description: "des3",
        isDone: false,
    }
])
