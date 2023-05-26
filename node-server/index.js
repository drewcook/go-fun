// A basic http server in node, which we will duplicate in golang

const express = require('express')
const app = express()
const port = 8000

app.use(express.json());
app.use(express.urlencoded({extended: true}));

// Sends back string data
app.get('/', (req, res) => {
  res.status(200).send("Welcome to this basic node server")
})

// Sends back JSON data
app.get('/get', (req, res) => {
	res.status(200).json({message: "Hello from dco.dev"})
})

// Accepts JSON data, returns it back
app.post('/post', (req, res) => {
	let myJson = req.body;
	res.status(200).send(myJson);
})

// Accepts form data, returns as JSON
app.post('/postform', (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
})


app.listen(port, () => {
  console.log(`Server is listening at http://localhost:${port}`)
})