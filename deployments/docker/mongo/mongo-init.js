db.createUser(
    {
        user: 'service',
        pwd: 'servicepassword',
        roles: [
            {
                role: 'readWrite',
                db: 'mongodb'
            }
        ]
    }
);

db.createCollection('users');
db.users.createIndex( { email: 1 }, { unique: true } );
