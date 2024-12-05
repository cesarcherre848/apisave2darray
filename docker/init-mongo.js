const dbName = "Condestable"; // Cambia por el nombre de tu base de datos
const collectionName = "2DData"; // Cambia por el nombre de tu colección

// Obtener una lista de las bases de datos existentes
const existingDbs = db.adminCommand({ listDatabases: 1 }).databases;

// Verificar si la base de datos ya existe
const dbExists = existingDbs.some(database => database.name === dbName);

if (!dbExists) {
    print(`Database '${dbName}' does not exist. Creating...`);
    const db = db.getSiblingDB(dbName);
    db.createCollection(collectionName);
    db[collectionName].insertOne({ name: "example", value: 42 }); // Insertar un documento de ejemplo
    print(`Database '${dbName}' and collection '${collectionName}' created successfully.`);
} else {
    print(`Database '${dbName}' already exists. Skipping creation.`);
}