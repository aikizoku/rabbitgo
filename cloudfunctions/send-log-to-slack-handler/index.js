var { IncomingWebhook } = require("@slack/client");
var moment = require("moment");

var LOG_COLORS = {
    DEBUG:    "#4175e1",
    INFO:     "#76a9fa",
    WARNING:  "warning",
    ERROR:    "danger",
    CRITICAL: "#ff0000",
};

var LINE_TIME_FORMAT = "YYYY-MM-DD HH:mm:ss.SSS";
var TIME_FORMAT      = "YYYY-MM-DD (ddd) HH:mm:ss";
var TIME_ZONE        = "+0900";

exports.main = (event, callback) => {
    var data = JSON.parse(new Buffer(event.data.data, "base64").toString());
    var payload = data.protoPayload;
    
    var appId = payload.appId.slice(payload.appId.indexOf("~") + 1);
    var logUrl = `https://console.developers.google.com/logs?project=${appId}`
        + `&service=appengine.googleapis.com&logName=appengine.googleapis.com%2Frequest_log&expandAll=true`
        + `&filters=request_id:${payload.requestId}`;
    
    var body = {
        attachments: [{
            title: `${payload.method} ${payload.resource} <${logUrl}|Open Log>`,
            text: (payload.line || []).map(l => 
                `\`${moment(l.time).utcOffset(TIME_ZONE).format(LINE_TIME_FORMAT)}\``
                + ` \`${l.severity}\``
                + ` ${l.logMessage}`).join("\n"),
            color: LOG_COLORS[data.severity],
            fields: [
                { title: "Level", value: data.severity, short: true },
                { title: "Timestamp", value: moment(data.timestamp).utcOffset(TIME_ZONE).format(TIME_FORMAT), short: true },
                { title: "Host", value: payload.host, short: true },
                { title: "Build Version", value: payload.versionId, short: true },
                { title: "HTTP Status", value: payload.status, short: true },
                { title: "Request IP", value: payload.ip, short: true },
                { title: "User Agent", value: payload.userAgent, short: false },
            ]
        }]
    };

    var webhook = new IncomingWebhook(process.env.SLACK_WEBHOOK_URL);
    webhook.send(body, (err, res) => {
        if (err) { console.error(err); }
        callback();
    });
};
