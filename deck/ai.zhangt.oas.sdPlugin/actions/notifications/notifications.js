const utoa = (data) => {
    return btoa(unescape(encodeURIComponent(data)));
};

const IconStats = {
    "teams": {
        "active": {
            "primary": "#5059C9",
            "secondary": "#7B83EB",
        },
        "inactive": {
            "primary": "#3C4048",
            "secondary": "#B2B2B2",
        },
        "error": {
            "primary": "#D21312",
            "secondary": "#F45050",
        },
    },
    "outlook": {
        "active": {
            "primary": "#0072c6",
            "secondary": "",
        },
        "inactive": {
            "primary": "#3C4048",
            "secondary": "",
        },
        "error": {
            "primary": "#D21312",
            "secondary": "",
        },
    }
}

const getStatusIcon = (appName, status) => {
    const colors = IconStats[appName][status]
    let svgString = ""
    if (appName == "outlook") {
        svgString = `<svg clip-rule="evenodd" fill-rule="evenodd" height="144" image-rendering="optimizeQuality" shape-rendering="geometricPrecision" text-rendering="geometricPrecision" viewBox="0 0 6876 6994" width="2457" xmlns="http://www.w3.org/2000/svg"><path d="M0 779L4033 0l-14 6994L0 6160zm1430 3632c-305-357-390-918-244-1384 203-648 718-867 1149-717 246 86 465 293 582 610 56 152 86 326 88 503 4 318-106 692-324 953-335 400-903 441-1250 35zm314-339c-150-223-191-573-120-864 99-404 352-541 563-447 121 54 228 183 285 381 27 95 42 203 43 314 2 198-52 432-159 595-164 250-442 275-612 22zm2552-2598h2341c131 0 238 107 238 238v86L5035 3039c-24 16-83 62-132 93-72 47-77 38-153-5-117-65-319-203-455-297V1474zm2580 875v2504c0 200-164 365-365 365H4296V3366c133 88 310 204 419 271 88 54 104 79 202 22 45-26 89-60 119-80l1840-1229z" fill="${colors.primary}"/></svg>`
    }
    if (appName == "teams") {
        svgString = `<?xml version="1.0" encoding="utf-8"?><svg version="1.1" id="Livello_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 2228.83 2073.33" style="enable-background:new 0 0 2228.83 2073.33" xml:space="preserve"><style type="text/css">.st0{fill:${colors.primary};} .st1{fill:${colors.secondary};} .st2{opacity:0.1;enable-background:new ;} .st3{opacity:0.2;enable-background:new ;} .st4{fill:url(#SVGID_1_);} .st5{fill:#FFFFFF;}</style><path class="st0" d="M1554.64,777.5h575.71c54.39,0,98.48,44.09,98.48,98.48l0,0v524.4c0,199.9-162.05,361.95-361.95,361.95l0,0	h-1.71c-199.9,0.03-361.97-162-362-361.9c0-0.02,0-0.03,0-0.05V828.97C1503.17,800.54,1526.21,777.5,1554.64,777.5L1554.64,777.5z"/><circle class="st0" cx="1943.75" cy="440.58" r="233.25"/><circle class="st1" cx="1218.08" cy="336.92" r="336.92"/><path class="st1" d="M1667.32,777.5H717.01c-53.74,1.33-96.26,45.93-95.01,99.68v598.1c-7.51,322.52,247.66,590.16,570.17,598.05	c322.51-7.89,577.67-275.53,570.17-598.05v-598.1C1763.58,823.43,1721.07,778.83,1667.32,777.5z"/><path class="st2" d="M1244,777.5v838.15c-0.26,38.44-23.55,72.96-59.09,87.6c-11.32,4.79-23.48,7.25-35.77,7.26H667.61	c-6.74-17.1-12.96-34.21-18.14-51.83c-18.14-59.48-27.4-121.31-27.47-183.49V877.02c-1.25-53.66,41.2-98.19,94.85-99.52H1244z"/><path class="st3" d="M1192.17,777.5v889.98c0,12.29-2.47,24.45-7.26,35.77c-14.63,35.54-49.16,58.83-87.6,59.09H691.97	c-8.81-17.1-17.1-34.21-24.36-51.83s-12.96-34.21-18.14-51.83c-18.14-59.48-27.4-121.31-27.47-183.49V877.02	c-1.25-53.66,41.2-98.19,94.85-99.52H1192.17z"/><path class="st3" d="M1192.17,777.5v786.31c-0.4,52.22-42.63,94.46-94.85,94.85H649.47c-18.14-59.48-27.4-121.31-27.47-183.49	V877.02c-1.25-53.66,41.2-98.19,94.85-99.52H1192.17z"/><path class="st3" d="M1140.33,777.5v786.31c-0.4,52.22-42.63,94.46-94.85,94.85H649.47c-18.14-59.48-27.4-121.31-27.47-183.49	V877.02c-1.25-53.66,41.2-98.19,94.85-99.52H1140.33z"/><path class="st2" d="M1244,509.52V672.8c-8.81,0.52-17.1,1.04-25.92,1.04s-17.1-0.52-25.92-1.04c-17.5-1.16-34.85-3.94-51.83-8.29	c-104.96-24.86-191.68-98.47-233.25-198c-7.15-16.71-12.71-34.07-16.59-51.83h258.65C1201.45,414.87,1243.8,457.22,1244,509.52z"/><path class="st3" d="M1192.17,561.35V672.8c-17.5-1.16-34.85-3.94-51.83-8.29c-104.96-24.86-191.68-98.47-233.25-198h190.23	C1149.62,466.7,1191.97,509.05,1192.17,561.35z"/><path class="st3" d="M1192.17,561.35V672.8c-17.5-1.16-34.85-3.94-51.83-8.29c-104.96-24.86-191.68-98.47-233.25-198h190.23	C1149.62,466.7,1191.97,509.05,1192.17,561.35z"/><path class="st3" d="M1140.33,561.35V664.5c-104.96-24.86-191.68-98.47-233.25-198h138.4	C1097.78,466.7,1140.13,509.05,1140.33,561.35z"/><linearGradient id="SVGID_1_" gradientUnits="userSpaceOnUse" x1="198.0988" y1="-1123.0724" x2="942.2333" y2="165.7381" gradientTransform="matrix(1 0 0 1 0 1515.3333)"><stop offset="0" style="stop-color:${colors.primary}"/><stop offset="0.5" style="stop-color:${colors.secondary}"/><stop offset="1" style="stop-color:${colors.secondary}"/></linearGradient><path class="st4" d="M95.01,466.5h950.31c52.47,0,95.01,42.54,95.01,95.01v950.31c0,52.47-42.54,95.01-95.01,95.01H95.01	c-52.47,0-95.01-42.54-95.01-95.01V561.51C0,509.04,42.54,466.5,95.01,466.5z"/><path class="st5" d="M820.21,828.19H630.24v517.3H509.21v-517.3H320.12V727.84h500.09V828.19z"/></svg>`
    }
    let b64 = `data:image/svg+xml;base64,${utoa(svgString)}`;
    return b64
}

const drawStateIcon = (ctx, appName, endpoint, fetchOption, unreadParser) => {
    fetch(endpoint, fetchOption)
        .then(response => response.json())
        .then(data => {
            const countName = appName + "UnreadCount";
            const count = unreadParser(data)
            if (window[countName] === undefined || count != window[countName]) {
                console.log(`[${countName}]Unread count changed from ${window[countName]} to ${count}`);
                $SD.websocket.send(JSON.stringify({
                    "event": "setTitle",
                    "context": ctx,
                    "payload": {
                        "title": count.toString(),
                        "target": 0
                    }
                }))
                let newIcon = count > 0 ? getStatusIcon(appName, "active") : getStatusIcon(appName, "inactive")
                $SD.websocket.send(JSON.stringify({
                    "event": "setImage",
                    "context": ctx,
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
                "context": ctx,
                "payload": {
                    "title": "err",
                    "target": 0
                }
            }))
            $SD.websocket.send(JSON.stringify({
                "event": "setImage",
                "context": ctx,
                "payload": {
                    "image": getStatusIcon(appName, "error"),
                    "target": 0
                }
            }))
        });
}

function unreadDisplay(jsonObj, appName, endpoint, fetchOption, unreadParser) {
    var jsn = jsonObj,
        context = jsonObj.context,
        displayTimer = 0,
        origContext = jsonObj.context,
        count = Math.floor(Math.random() * Math.floor(10));

    function createDisplay() {
        if (displayTimer === 0) {
            displayTimer = setInterval(function (sx) {
                drawStateIcon(context, appName, endpoint, fetchOption, unreadParser)
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

const generateUnreadAction = (actionId, appName, tabId, endpoint, fetchOption, unreadParser) => {
    return {
        type: actionId,
        cache: {},

        onWillAppear: function (jsn) {
            drawStateIcon(jsn.context, appName, endpoint, fetchOption, unreadParser)
            const display = new unreadDisplay(jsn, appName, endpoint, fetchOption, unreadParser);
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
            fetch(`http://127.0.0.1:8090/api/apps/chromium/tabs/${tabId}/activate`, {
                method: "POST"
            })
        }
    }
}
