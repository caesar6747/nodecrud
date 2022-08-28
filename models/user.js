const Sequelize = require("sequelize");
const db = require("../config/db");

const user = db.define(
    "user",
    {
        id: {type: Sequelize.STRING,
            allowNull: false,
            primaryKey: true},
        username: {type: Sequelize.STRING,
            allowNull: false},
        password: {type: Sequelize.STRING,
            allowNull: false},
    },
    {
        freezeTableName: true,
        timestamps: false
    }
);

module.exports = user;