const dateTimeApiBaseUrl = 'http://35.204.56.198:8080';
const quoteApiBaseUrl = 'http://34.90.216.212:8080';

function getDateTime() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', dateTimeApiBaseUrl + '/datetime');
    xhr.onreadystatechange = function() {
        console.log(xhr.responseText);
        var data;
        if (xhr.readyState == 4 && xhr.status == 200) {
            data = JSON.parse(xhr.responseText);
            document.getElementById('currentDateTime').innerText = data.dateTime;
        }
    }
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send();
}

function getQuote() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', quoteApiBaseUrl + '/quote');
    xhr.onreadystatechange = function() {
        console.log(xhr.responseText);
        var data;
        if (xhr.readyState == 4 && xhr.status == 200) {
            data = JSON.parse(xhr.responseText);
            document.getElementById('currentQuote').innerText = data.quote;
        }
    }
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send();
}
