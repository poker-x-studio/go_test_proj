'use strict';

const crypto = require('crypto');

const TAG = '[node]';

function signRequest(clientId, clientSecret, uri, requestMethod, body) {
    console.log("clientId:", clientId)
    console.log("clientSecret:", clientSecret)
    console.log("uri:", uri)
    console.log("requestMethod:", requestMethod)
    console.log("body:", body)

    const requestTimeStamp = Math.round(Date.now() / 1000);
    const requestUri = encodeURIComponent(uri.toLowerCase());
    console.log("requestUri:", requestUri)

    let requestContentBase64String = '';
    if (body !== undefined && body !== null && body.trim() !== '') {
        requestContentBase64String = crypto.createHash('md5').update(body).digest("base64")
    }
    console.log("requestContentBase64String:", requestContentBase64String)

    const nonce = crypto.randomBytes(16).toString("hex");
    const signatureRawData = clientId + requestMethod + requestUri + requestTimeStamp + nonce + requestContentBase64String;
    console.log('signatureRawData:', signatureRawData);

    const signature = crypto.createHmac('SHA256', clientSecret).update(signatureRawData).digest('base64');
    console.log('signature:', signature);

    return clientId + ':' + signature + ':' + nonce + ':' + requestTimeStamp;
}

module.exports = {
	signRequest,
};

console.log('process.argv:', process.argv);
var sign = signRequest(process.argv[2],process.argv[3],process.argv[4],process.argv[5],process.argv[6]);
console.log('sign:', sign);
