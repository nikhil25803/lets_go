const express = require("express")
const app = express()
const port = 3000


app.use(express.json())
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
    res.status(200).json({
        message: "Hello this is test server"
    })
})

app.get('/get', (req, res) => {
    res.status(200)
        .json({
            message: "You made a get request successfully"
        })
})

app.post('/post', (req, res) => {
    res.status(200)
        .json({
            message: req.body
        })
})

app.post('/postForm', (req, res) => {
    res.status(200)
        .send(JSON.stringify(req.body))
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})