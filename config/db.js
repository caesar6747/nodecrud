const sequelize = require("sequelize");
const seq = require("sequelize");

const db = new sequelize("dbcrud", "root", "Alwaysopen1", {
    dialect: "mysql"
});

db.sync({});

module.exports = db;
