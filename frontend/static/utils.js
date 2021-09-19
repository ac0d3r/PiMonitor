/*
    request utils
*/

function getHTTPObject() {
    if (typeof XMLHttpRequest != 'undefined') {
        return new XMLHttpRequest();
    } try {
        return new ActiveXObject("Msxml2.XMLHTTP");
    } catch (e) {
        try {
            return new ActiveXObject("Microsoft.XMLHTTP");
        } catch (e) { }
    }
    return null;
}

function getUrl(uri) {
    return "http://" + document.location.host + uri;
}

function sendJSON(xmlhttp, data) {
    xmlhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xmlhttp.send(JSON.stringify(data));
}

function carryUserToken(xmlhttp) {
    xmlhttp.setRequestHeader("Authorization", "Bearer " + getCookieValueByKey("token"));
}

/*
    cookie && user token
*/

function getCookieValueByKey(key) {
    return decodeURIComponent(document.cookie.replace(new RegExp("(?:(?:^|.*;)\\s*" + encodeURIComponent(key).replace(/[-.+*]/g, "\\$&") + "\\s*\\=\\s*([^;]*).*$)|^.*$"), "$1")) || null;
}

function saveUserToken(data) {
    document.cookie = "id=" + data.id;
    document.cookie = "token=" + data.token;
    document.cookie = "username=" + data.username;
}

function removeUserToken() {
    document.cookie = "id=";
    document.cookie = "token=";
    document.cookie = "username=";
}
