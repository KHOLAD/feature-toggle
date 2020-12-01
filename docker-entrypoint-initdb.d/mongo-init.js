db.auth('admin', 'password');
db = db.getSiblingDB('featureTable');

db.createCollection("customers");
db.createCollection("features");

db.runCommand(
    {
        insert: "customers",
        documents: [
            {
                _id: "1",
                name: "Swiss",
                features: []
            },
            {
                _id: "2",
                name: "C.T.co",
                features: []
            }
        ]
    }
)

db.createUser({
    user: 'featureAdmin',
    pwd: 'password',
    roles: [
        {
            role: 'root',
            db: 'admin',
        },
    ],
});
