
const express = require('express'); 
const axios = require('axios')
  
const app = express(); 
const PORT = 3001; 
  

app.get('/', async (req, res) => {
    const response = await axios.get('https://jsonplaceholder.typicode.com/posts');
    res.status(200)
    res.send(response.data)
})

app.get('/cpu-intensive', async (req,res) => {
    const start = new Date();
    let sum = 0;
    for (let i = 0; i < 100000000; i++) {
        sum += i;
    }
    const elapsed = new Date().getTime() - start.getTime();
    res.send(`CPU-intensive task completed in ${elapsed} milliseconds\n`);
})

app.listen(PORT, (error) =>{ 
    if(!error) 
        console.log("Server is Successfully Running, and App is listening on port "+ PORT) 
    else 
        console.log("Error occurred, server can't start", error); 
    } 
); 