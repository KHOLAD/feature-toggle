db.auth('admin', 'password');
db = db.getSiblingDB('featureTable');

const collections = ['customers', 'features'];

collections.forEach((col) => {
    db.createCollection(col)
})

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

db.runCommand(
    {
        insert: "customers",
        documents: [
            {
                name: "SwissCo",
                features: []
            },
            {
                name: "C.T.co",
                features: []
            }
        ]
    }
)
