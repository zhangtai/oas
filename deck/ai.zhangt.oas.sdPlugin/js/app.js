/* global $SD */
$SD.on('connected', conn => connected(conn));

const teamsAction = generateUnreadAction(
    'ai.zhangt.oas.teams.action',
    "teams",
    1,
    "http://localhost:8090/api/apps/chromium/tabs/1",
    { method: "GET" },
    (data) => { let m = data.title.match(/^\((?<count>\d+)\)/); return m != null ? m.groups.count : 0 }
)

const outlookAction = generateUnreadAction(
    'ai.zhangt.oas.outlook.action',
    "outlook",
    2,
    "http://localhost:8090/api/apps/chromium/tabs/2/execute",
    {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ script: "document.querySelector('._n_B4._n_z4 span[autoid=_n_41]')?.textContent || 0" })
    },
    (data) => parseInt(data)
)

const upcommingremindersAction = geneartearemindersAction("ai.zhangt.oas.reminders.action")

function connected(jsn) {
    console.log('Connected Plugin:', jsn);

    /** Teams */
    $SD.on('ai.zhangt.oas.teams.action.willAppear', jsonObj =>
        teamsAction.onWillAppear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.teams.action.willDisappear', jsonObj =>
        teamsAction.onWillDisappear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.teams.action.keyUp', jsonObj =>
        teamsAction.onKeyUp(jsonObj)
    );

    /** Outlook */
    $SD.on('ai.zhangt.oas.outlook.action.willAppear', jsonObj =>
        outlookAction.onWillAppear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.outlook.action.willDisappear', jsonObj =>
        outlookAction.onWillDisappear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.outlook.action.keyUp', jsonObj =>
        outlookAction.onKeyUp(jsonObj)
    );

    /** Event Countdown */
    $SD.on('ai.zhangt.oas.reminders.action.willAppear', jsonObj =>
        upcommingremindersAction.onWillAppear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.reminders.action.willDisappear', jsonObj =>
        upcommingremindersAction.onWillDisappear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.reminders.action.keyUp', jsonObj =>
        upcommingremindersAction.onKeyUp(jsonObj)
    );
}
