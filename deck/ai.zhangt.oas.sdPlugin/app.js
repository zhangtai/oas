/* global $SD */
$SD.on('connected', conn => connected(conn));

function connected(jsn) {
    console.log('Connected Plugin:', jsn);

    /** Notifications */
    $SD.on('ai.zhangt.oas.notifications.action.willAppear', jsonObj =>
        notificationsAction.onWillAppear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.notifications.action.willDisappear', jsonObj =>
        notificationsAction.onWillDisappear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.notifications.action.keyUp', jsonObj =>
        notificationsAction.onKeyUp(jsonObj)
    );

    /** Event Countdown */
    $SD.on('ai.zhangt.oas.reminders.action.willAppear', jsonObj =>
        nextReminderAction.onWillAppear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.reminders.action.willDisappear', jsonObj =>
        nextReminderAction.onWillDisappear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.reminders.action.keyUp', jsonObj =>
        nextReminderAction.onKeyUp(jsonObj)
    );

    /** Event Countdown */
    $SD.on('ai.zhangt.oas.charts.action.willAppear', jsonObj =>
        chartsAction.onWillAppear(jsonObj)
    );
    $SD.on('ai.zhangt.oas.charts.action.willDisappear', jsonObj =>
        chartsAction.onWillDisappear(jsonObj)
    );
}
