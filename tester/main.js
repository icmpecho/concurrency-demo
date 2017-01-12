const request = require('request')
const n = process.argv[2] || 1
const start = (new Date).getTime()

for(let i=0; i < n; i++) {
  request('http://localhost:8000', (error, response, body) => {
    const elapsed = (new Date).getTime() - start
    console.log(`Received response ${i+1} at ${elapsed}ms: ${body}`)
  })
}

