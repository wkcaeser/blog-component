function uuid(len, radix) {
    var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    var uuid = [], i;
    radix = radix || chars.length;

    if (len) {
        // Compact form
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
    } else {
        // rfc4122, version 4 form
        var r;

        // rfc4122 requires these characters
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';

        // Fill in random data.  At i==19 set the high bits of clock sequence as
        // per rfc4122, sec. 4.1.5
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }

    return uuid.join('');
}
// 这个可以指定长度和基数。比如

// "098F4D35"

function buildUuidInfo() {
    return {
        uuid: uuid(32, 16),
        expireTime: new Date().getTime() + 30 * 24 * 3600 * 1000
    }
}

function initUuidInfo(uuidInfoStr) {
    var uuidInfo;
    if (!uuidInfoStr) {
        uuidInfo = buildUuidInfo();
    } else {
        try {
            uuidInfo = JSON.parse(uuidInfoStr);
            if (uuidInfo.expireTime < new Date().getTime()) {
                uuidInfo = buildUuidInfo();
            }
        } catch (error) {
            console.error(error);
            uuidInfo = buildUuidInfo();
        }
    }

    if ((!uuidInfo.uuid) || (!uuidInfo.expireTime)) {
        uuidInfo = buildUuidInfo();
    }

    console.log(uuidInfo);
    return JSON.stringify(uuidInfo);
}

function initStorageMetric() {
    var storage = window.localStorage;
    var uuidInfoStr = storage.getItem("browerUuidInfo");

    storage.setItem("browerUuidInfo", initUuidInfo(uuidInfoStr));
}

function initCookieMetric() {
    var uuidInfoStr = document.cookie.replace(/(?:(?:^|.*;\s*)browerUuidInfo\s*\=\s*([^;]*).*$)|^.*$/, "$1");

    document.cookie = "browerUuidInfo=" + initUuidInfo(uuidInfoStr) + "; expires=Thu, 18 Dec 2043 12:00:00 GMT;";
}

function initUuid() {
    if (window.localStorage) {
        initStorageMetric();
    } else {
        initCookieMetric();
    }
}

function getUuid() {
    var uuidInfoStr = window.localStorage.getItem("browerUuidInfo");

    if (!uuidInfoStr) {
        uuidInfoCookieStr = document.cookie.replace(/(?:(?:^|.*;\s*)browerUuidInfo\s*\=\s*([^;]*).*$)|^.*$/, "$1");
    }

    var uuidInfo = JSON.parse(uuidInfoStr);
    return uuidInfo.uuid;
}

function sendMetric() {
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                console.log(xhr.responseText);
                var data = JSON.parse(xhr.responseText);
                document.getElementById("peopleTotal").innerHTML = data.peopleTotal;
                document.getElementById("peoplePv").innerHTML = data.peoplePv;
            }
        }
    };

    xhr.open("get", "https://tapi.wkcaeser.com/o/info?biz=pv&uuid=" + getUuid() + "&" + "uri=" + document.URL, true);
    xhr.send(null);
}


initUuid();
console.log(getUuid());

sendMetric();
