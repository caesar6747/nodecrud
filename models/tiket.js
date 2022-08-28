const Sequelize = require("sequelize");
const db = require("../config/db");

const tiket = db.define(
    "tiket",
    {
        id: {type: Sequelize.STRING,
            allowNull: false,
            primaryKey: true},
        nama: {type: Sequelize.STRING,
            allowNull: false},
        layanan: {type: Sequelize.STRING,
            allowNull: false},
        tgl: {type: Sequelize.DATE,
            allowNull: false},
    },
    {
        freezeTableName: true,
        timestamps: false 
    }
);

module.exports = tiket;