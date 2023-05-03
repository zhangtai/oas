const utoa = (data) => {
    return btoa(unescape(encodeURIComponent(data)));
};

const getStatusIcon = (svgTemplate, primaryColor, secondaryColor) => {
    let svgString = svgTemplate;
    if (primaryColor) {
        svgString = svgString.replaceAll("OAS_SVG_PRIMARYCOLOR", primaryColor)
    }
    if (secondaryColor) {
        svgString = svgString.replaceAll("OAS_SVG_SECONDARYCOLOR", secondaryColor)
    }
    let b64 = `data:image/svg+xml;base64,${utoa(svgString)}`;
    return b64
}

const drawStateIcon = (jsn) => {
    const settings = jsn.payload.settings;
    const parser = eval(settings.parserFunction)
    // console.log(settings)
    fetchOption = {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            script: settings.fetchPostJs,
            urlStartsWith: settings.urlStartsWith
        })
    }
    // console.log(fetchOption)
    fetch("http://localhost:8090/api/apps/chromium/tab", fetchOption)
        .then(response => response.json())
        .then(data => {
            // console.log(data)
            const countName = jsn.context + "UnreadCount";
            const count = parser(data)
            if (window[countName] === undefined || count != window[countName]) {
                console.log(`[${countName}]Unread count changed from ${window[countName]} to ${count}`);
                $SD.websocket.send(JSON.stringify({
                    "event": "setTitle",
                    "context": jsn.context,
                    "payload": {
                        "title": count.toString(),
                        "target": 0
                    }
                }))
                let newIcon = count > 0 
                    ? getStatusIcon(settings.svgTemplate, settings.svgColorActivePrimary, settings.svgColorActiveSecondary)
                    : getStatusIcon(settings.svgTemplate, settings.svgColorInactivePrimary, settings.svgColorInactiveSecondary)
                $SD.websocket.send(JSON.stringify({
                    "event": "setImage",
                    "context": jsn.context,
                    "payload": {
                        "image": newIcon,
                        "target": 0
                    }
                }))
            }
            window[countName] = count;
        })
        .catch(error => {
            console.error('There was an error!', error);
            $SD.websocket.send(JSON.stringify({
                "event": "setTitle",
                "context": jsn.context,
                "payload": {
                    "title": "err",
                    "target": 0
                }
            }))
            $SD.websocket.send(JSON.stringify({
                "event": "setImage",
                "context": jsn.context,
                "payload": {
                    "image": getStatusIcon(settings.svgTemplate, settings.svgColorErrorPrimary, settings.svgColorErrorSecondary),
                    "target": 0
                }
            }))
        });
}

function unreadDisplay(jsonObj) {
    var jsn = jsonObj,
        context = jsonObj.context,
        displayTimer = 0,
        origContext = jsonObj.context,
        count = Math.floor(Math.random() * Math.floor(10));

    function createDisplay() {
        if (displayTimer === 0) {
            displayTimer = setInterval(function (sx) {
                drawStateIcon(jsn)
                count++;
            }, 15000);
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

const notificationsAction = {
    type: 'ai.zhangt.oas.notifications.action',
    cache: {},

    onWillAppear: function (jsn) {
        drawStateIcon(jsn)
        const display = new unreadDisplay(jsn);
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
        fetchOption = {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "script": `activate application "${jsn.payload.settings.appName}"`
            })
        }
        fetch("http://localhost:8090/api/system/applescript", fetchOption)
    }
}
