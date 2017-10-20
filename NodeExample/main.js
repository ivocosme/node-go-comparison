const app = require('express')();
const law = 'Don’t worry if it doesn’t work right. If everything did, you’d be out of a job. - (Mosher’s Law of Software Engineering)\n';

app.get('/blocking-eventloop', (req, res) => {
  var counter = 0;
  const endDate = new Date(Date.now() + (5000));
  do {
    counter += 10;
  } while (Date.now() < endDate);
  res.send(law);
});

app.get('/non-blocking-eventloop', (req, res) => {
  setTimeout(() => res.send(law), 5000);
});

app.listen(3000, () => console.log('App listening on port 3000'));