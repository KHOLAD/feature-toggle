db.auth('admin', 'password')
db = db.getSiblingDB('featureTable')
db.createCollection("features");

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
