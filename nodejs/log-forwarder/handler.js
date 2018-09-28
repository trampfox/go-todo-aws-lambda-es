/*
 * Sample node.js code for AWS Lambda to upload the JSON documents
 * pushed from CloudWatch to Amazon Elasticsearch.
 */

/* == Imports == */
var AWS = require('aws-sdk');
var path = require('path');
var zlib = require('zlib');

/* == Globals == */
var esDomain = {
    region: process.env.REGION,
    endpoint: process.env.ENDPOINT,
    index: process.env.INDEX,
    doctype: process.env.DOC_TYPE
};
var endpoint = new AWS.Endpoint(esDomain.endpoint);
/*
 * The AWS credentials are picked up from the environment.
 * They belong to the IAM role assigned to the Lambda function.
 * Since the ES requests are signed using these credentials,
 * make sure to apply a policy that allows ES domain operations
 * to the role.
 */
var creds = new AWS.EnvironmentCredentials('AWS');


/* Lambda "main": Execution begins here */
exports.handler = function(event, context) {
    console.log(JSON.stringify(event, null, '  '));
    
    var payload = new Buffer(event.awslogs.data, 'base64');
    zlib.gunzip(payload, function(e, result) {
        if (e) { 
            context.fail(e);
        } else {
            result = JSON.parse(result.toString('ascii'));
            console.log("Event Data:", JSON.stringify(result, null, 2));
            postToES(JSON.stringify(result, null, 2), context);
        }
    });
}


/*
 * Post the given document to Elasticsearch
 */
function postToES(doc, context) {
    console.log("Posting document...");
    console.log(doc);
    var req = new AWS.HttpRequest(endpoint);

    req.method = 'POST';
    req.path = path.join('/', esDomain.index, esDomain.doctype);
    req.region = esDomain.region;
    req.headers['presigned-expires'] = false;
    req.headers['Host'] = endpoint.host;
    req.body = doc;

    var signer = new AWS.Signers.V4(req , 'es');  // es: service code
    signer.addAuthorization(creds, new Date());

    var send = new AWS.NodeHttpClient();
    send.handleRequest(req, null, function(httpResp) {
        var respBody = '';
        httpResp.on('data', function (chunk) {
            respBody += chunk;
        });
        httpResp.on('end', function (chunk) {
            console.log('Response: ' + respBody);
            context.succeed('Lambda added document ' + doc);
        });
    }, function(err) {
        console.log('Error: ' + err);
        context.fail('Lambda failed with error ' + err);
    });
}