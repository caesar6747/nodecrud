const express = require("express");
const mysql = require("mysql");
const app = express();
const port = 8080;

const connection = mysql.createConnection({
    host: "localhost",
    user: "root",
    password: "Alwaysopen1",
    database: "dbcrud"
});

connection.connect(function(error){
    if(!!error){
        console.log("error");
    }
    else{
        console.log("connected");
    }
});
app.get("/", function(req, res){
    connection.query("SELECT *FROM user", function(error, rows, fields){
        if(!!error){
            console.log("error in the query");
        }
        else{
            console.log("Successful query");
        }
    });
});
app.listen(1336);