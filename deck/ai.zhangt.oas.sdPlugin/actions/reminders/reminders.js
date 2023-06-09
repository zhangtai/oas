const toHHMM = (seconds) => {
    overdueSign = ""
    if (seconds < 0) {
        seconds = 0 - seconds
        overdueSign = "-"
    }
    let hours = Math.floor(seconds / 3600);
    let minutes = Math.floor((seconds - (hours * 3600)) / 60);
    if (hours < 1) { hours = "0" + hours; }
    if (minutes < 1) { minutes = "0" + minutes; }
    return overdueSign + hours.toString().padStart(2, '0') + ':' + minutes.toString().padStart(2, '0');
}

var latestReminderUrl = "";

const drawCountdownIcon = (ctx) => {
    fetch("http://localhost:8090/api/apps/reminders/all", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            dueDate: (new Date()).toISOString(),
        })
    })
        .then(resp => resp.json())
        .then(data => {
            console.log(data)
            let leftTime = "--:--";
            let title = "NA";
            let color = "#8cac8c";
            let host = "Me";
            let timeFontSize = 50;
            let hostStartPosition = 42;
            if (data) {
                reminder = data[0]
                const eventTime = new Date(reminder.dueDate)
                console.log(eventTime)
                const seconds = (eventTime - (new Date())) / 1000
                if (seconds < 5 * 60) {
                    color = "red";
                }
                leftTime = toHHMM(seconds);
                if (leftTime.length > 5) {
                    timeFontSize = 44;
                }
                title = reminder.title;
                if (reminder.meta?.host) {
                    host = reminder.meta.host;
                    hostStartPosition = 72 - host.length * 8;
                }
                if (reminder.meta?.url) {
                    latestReminderUrl = reminder.meta.url;
                }
            }
            let svgString = `<svg height="144px" width="144px" xmlns="http://www.w3.org/2000/svg"><text x="${hostStartPosition}" y="40" style="fill: #c27fcd; font-size: 36px;">${host}</text><text x="6" y="94" style="fill: ${color}; font-size: ${timeFontSize}px;">${leftTime}</text></svg>`
            $SD.websocket.send(JSON.stringify({
                "event": "setImage",
                "context": ctx,
                "payload": {
                    "image": `data:image/svg+xml;base64,${utoa(svgString)}`,
                    "target": 0
                }
            }))
            $SD.websocket.send(JSON.stringify({
                "event": "setTitle",
                "context": ctx,
                "payload": {
                    "title": title,
                    "target": 0
                }
            }))
        })
}

function countdownDisplay(jsonObj) {
    var jsn = jsonObj,
        context = jsonObj.context,
        displayTimer = 0,
        origContext = jsonObj.context,
        count = Math.floor(Math.random() * Math.floor(10));

    function createDisplay() {
        if (displayTimer === 0) {
            displayTimer = setInterval(function (sx) {
                drawCountdownIcon(context)
                count++;
            }, 30000);
        } else {
            window.clearInterval(displayTimer);
            displayTimer = 0;
        }
    }

    function destroyDisplay() {
        if (displayTimer !== 0) {
            window.clearInterval(displayTimer);
            displayTimer = 0;
        }
    }

    createDisplay();

    return {
        displayTimer: displayTimer,
        origContext: origContext,
        destroyDisplay: destroyDisplay,
    };
}

const nextReminderAction = {
    type: "ai.zhangt.oas.reminders.action",
    cache: {},

    onWillAppear: function (jsn) {
        drawCountdownIcon(jsn.context)
        const display = new countdownDisplay(jsn);
        this.cache[jsn.context] = display;
    },

    onWillDisappear: function (jsn) {
        let found = this.cache[jsn.context];
        if (found) {
            found.destroyDisplay();
            delete this.cache[jsn.context];
        }
    },

    onKeyUp: function (jsn) {
        fetch("http://localhost:8090/api/system/command", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                command: ["open", latestReminderUrl]
            })
        })
    }
}
