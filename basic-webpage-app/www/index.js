//const apiBaseUrl = 'http://datetime.default.svc.cluster.local:8080';
const apiBaseUrl = 'http://35.204.56.198:8080';

function getDateTime() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', apiBaseUrl + '/datetime');
    xhr.onreadystatechange = function() {
        console.log(xhr.responseText);
        var data;
        var isPrime;
        if (xhr.readyState == 4 && xhr.status == 200) {
            data = JSON.parse(xhr.responseText);
            document.getElementById('currentDateTime').innerText = data.dateTime;
        }
    }
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send();
}
