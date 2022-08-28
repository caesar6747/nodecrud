const express = require("express");
const bp = require("body-parser")
const app = express();
const port = 1333;

const db = require("./config/db");
const tiket = require("./models/tiket");
const user = require("./models/user");

app.use(bp.json())
app.use(bp.urlencoded({extended:true}))

app.get("/", (req, res) => {
    res.send("Hello JS");
    console.log(req);
    console.log(user);
})

app.get("/read", async (req, res) => {
    try{
        const getTiket = await tiket.findAll();
        res.json(getTiket);
        console.log("start");
        console.log(getTiket.every(getTiket => getTiket instanceof tiket));
        console.log("end");
    }
    catch(err){
        console.log(err.message);
        console.log("error ternyata")
    }
})

app.post("/update", async (req, res) => {
    try {
        console.log("perintah update")
        console.log(req.body)
        const {id, nama, layanan, tgl} = req.body;
        const updateTiket = await tiket.update({
            layanan, tgl
        },{where: {id : id}}
        )
        await updateTiket;
        console.log("data diupdate")
        //console.log(newUser)
        res.json(updateTiket);
    } catch (error) {
        console.error(error.message)
    }
})

app.delete("/del/:id", async (req, res) => {
    try {
        console.log("perintah delete")
        const id = req.params.id;
        const deleteTiket = await tiket.destroy({
            where: {id : id}
        })
        await deleteTiket;
        console.log("data ditambahkan")
        //console.log(newUser)
        res.json(deleteTiket);
    } catch (error) {
        console.error(error.message)
    }
})

app.post("/add", async (req, res) => {
    try{
        const {id, nama, layanan, tgl} = req.body;
        const newTiket = new tiket({
            id, nama, layanan, tgl
        })
        console.log(layanan)
        console.log(req.body)
        await newTiket.save();
        console.log("data ditambahkan")
        //console.log(newUser)
        res.status.json(newTiket);
    }
    catch(err){
        console.error(err.message);
        res.status(500).send("server error")
    }
})

app.get("/user", async (req, res) =>{
    try{
        const getAllUser = await user.findAll();
        res.json(getAllUser);
        console.log("start");
        console.log(getAllUser.every(getAllUser => getAllUser instanceof user));
        console.log("end");
    }
    catch(err){
        console.log(err.message);
        console.log("error ternyata")
    }
})

db.authenticate().then(() => 
    console.log("berhasil")
);

app.listen(port, () => {
    console.log("complete ..");
})