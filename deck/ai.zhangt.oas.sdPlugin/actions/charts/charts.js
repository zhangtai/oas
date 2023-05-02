const randomIntFromInterval = (min, max) => { // min and max included 
  return Math.floor(Math.random() * (max - min + 1) + min)
}

const drawChartIcon = (ctx) => {
  now = (new Date()).getTime()
  let data = [
    [now - 86400000 * 5, randomIntFromInterval(5, 20)],
    [now - 86400000 * 4, randomIntFromInterval(5, 20)],
    [now - 86400000 * 3, randomIntFromInterval(5, 20)],
    [now - 86400000 * 2, randomIntFromInterval(5, 20)],
    [now - 86400000 * 1, randomIntFromInterval(5, 20)],
    [now , randomIntFromInterval(5, 20)]
  ]
  let options = {
    xaxis: {
      type: 'datetime',
      labels: { show: false },
    },
    yaxis: {
      min: 0,
      max: 25,
      labels: { show: false },
    },
    series: [
      {
        name: "usage",
        data: data
      }
    ],
    chart: {
      height: 144,
      width: 144,
      type: 'area',
      toolbar: { show: false },
      animations: { enabled: false },
      zoom: { enabled: false },
      sparkline: { enabled: true, },
    },
    fill: {
      type: "solid",
      opacity: 1,
      colors: [ data[data.length-1][1] > 12 ? "#D14D72" : "#B0DAFF" ]
    },
    dataLabels: { enabled: false },
    tooltip: { enabled: false },
    stroke: { curve: 'smooth' },
    grid: { show: false },
    legend: { show: false },
  };

  if (!document.querySelector(`#chart-${ctx}`)) {
    let chartDiv = document.createElement("div");
    chartDiv.id = `chart-${ctx}`
    document.body.prepend(chartDiv)
  }

  let chart = new ApexCharts(document.querySelector(`#chart-${ctx}`), options);
  chart.render()
  $SD.websocket.send(JSON.stringify({
    "event": "setImage",
    "context": ctx,
    "payload": {
      "image": `data:image/svg+xml;base64,${utoa(chart.paper().svg().replaceAll(/rgba\((\d+),(\d+),(\d+),(\d|\d\.\d)\)/g, "rgb($1,$2,$3)"))}`,
      "target": 0
    }
  }))
}

function chartDisplay(jsonObj) {
  var jsn = jsonObj,
    context = jsonObj.context,
    displayTimer = 0,
    origContext = jsonObj.context,
    count = Math.floor(Math.random() * Math.floor(10));

  function createDisplay() {
    if (displayTimer === 0) {
      displayTimer = setInterval(function (sx) {
        drawChartIcon(context)
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

const chartsAction = {
  type: "ai.zhangt.oas.charts.action",
  cache: {},

  onWillAppear: function (jsn) {
    drawChartIcon(jsn.context)
    const display = new chartDisplay(jsn);
    this.cache[jsn.context] = display;
  },

  onWillDisappear: function (jsn) {
    let found = this.cache[jsn.context];
    if (found) {
      found.destroyDisplay();
      delete this.cache[jsn.context];
    }
  },
}
