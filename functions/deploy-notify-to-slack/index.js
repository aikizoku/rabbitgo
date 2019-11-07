const { IncomingWebhook } = require('@slack/webhook');
var moment = require("moment");

var COLORS = {
  SUCCESS:        "#00FF40",
  FAILURE:        "warning",
  INTERNAL_ERROR: "warning",
  TIMEOUT:        "warning",
};
var TIME_FORMAT      = "YYYY-MM-DD HH:mm:ss";
var TIME_ZONE        = "+0900";

const url = process.env.SLACK_WEBHOOK_URL;
const webhook = new IncomingWebhook(url);

// subscribeSlack is the main function called by Cloud Functions.
module.exports.subscribeSlack = (pubSubEvent, context) => {
  const build = eventToBuild(pubSubEvent.data);

  const status = ['SUCCESS', 'FAILURE', 'INTERNAL_ERROR', 'TIMEOUT'];
  if (status.indexOf(build.status) === -1 || !build.substitutions._ZEUS_SERVICE_ID) {
    return;
  }

  // Send message to Slack.
  const message = createSlackMessage(build);
  webhook.send(message);
};

// eventToBuild transforms pubsub event message to a build object.
const eventToBuild = (data) => {
  return JSON.parse(Buffer.from(data, 'base64').toString());
}

// createSlackMessage creates a message from a build object.
const createSlackMessage = (build) => {
  const message = {
    attachments: [{
      title: `Build Result`,
      title_link: build.logUrl,
      color: COLORS[build.status],
      fields: [
        { title: 'Status', value: build.status, short: true },
        { title: 'Finish Time', value: moment(build.finishTime).utcOffset(TIME_ZONE).format(TIME_FORMAT), short: true },
        { title: 'Project', value: build.projectId, short: true },
        { title: 'Service', value: build.substitutions._ZEUS_SERVICE_ID, short: true },
        { title: 'Build ID', value: build.id, short: false },
      ]
    }]
  };
  return message;
}
