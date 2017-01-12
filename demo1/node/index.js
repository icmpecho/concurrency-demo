const express = require('express')
const faker = require('faker')
const app = express()

const query_data = (cb) => {
  setTimeout(() => {
    data = faker.name.findName()
    cb(data)
  }, 100)
}

app.get('/', (req, res) => {
  query_data((data) => {
    res.send({data: data})
  })
})

app.listen(8000, () => {
  console.log('App is listening on port 8000')
})
