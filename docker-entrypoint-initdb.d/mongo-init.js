db.auth('admin-user', 'password')

db = db.getSiblingDB('featureTable')

db.createUser({
    user: 'admin',
    pwd: 'password',
    roles: [
        {
            role: 'root',
            db: 'admin',
        },
    ],
});

db.featureTable.insert({})