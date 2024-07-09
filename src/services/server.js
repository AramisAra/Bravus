// backend/server.js
const express = require('express');
const { google } = require('googleapis');
const path = require('path');
const app = express();
const PORT = 5000;

const sheets = google.sheets('v4');

app.get('/api/sheets', async (req, res) => {
  try {
    const auth = new google.auth.GoogleAuth({
      keyFile: path.join(__dirname, 'path/to/your/service-account-file.json'), // Update the path
      scopes: ['https://www.googleapis.com/auth/spreadsheets.readonly'],
    });

    const client = await auth.getClient();
    const spreadsheetId = 'your_spreadsheet_id'; // Replace with your Spreadsheet ID

    const response = await sheets.spreadsheets.values.get({
      auth: client,
      spreadsheetId,
      range: 'Sheet1!A1:D10', // Adjust the range as per your need
    });

    res.json(response.data.values);
  } catch (error) {
    console.error('Error fetching Google Sheets data', error);
    res.status(500).send('Internal Server Error');
  }
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
