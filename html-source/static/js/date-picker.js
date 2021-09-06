(function () {
    'use strict'
    const elem = document.getElementById('date-picker');
    if (elem) {
        const rangepicker = new DateRangePicker(elem, {
            format: "yyyy-mm-dd"
        });
    } else {
        console.log('hey')
    }
})()