(function() {
    let attention = Prompt()
    function Prompt() {
        let toast = function (c) {
            const {
                icon = "success",
                title = "",
                position = "top-end"
            } = c;
            const Toast = Swal.mixin({
                toast: true,
                icon: icon,
                title: title,
                position: position,
                showConfirmButton: false,
                timer: 3000,
                timerProgressBar: true,
                didOpen: (toast) => {
                    toast.addEventListener('mouseenter', Swal.stopTimer)
                    toast.addEventListener('mouseleave', Swal.resumeTimer)
                }
            })

            Toast.fire()
        }
        let success = function (c) {
            const {
                icon = 'success',
                title = '',
                text = '',
            } = c;
            Swal.fire({
                icon: icon,
                title: title,
                text: text
            })
        }
        let error = function (c) {
            const {
                icon = 'error',
                title = '',
                text = '',
            } = c;
            Swal.fire({
                icon: icon,
                title: title,
                text: text
            })
        }
        let custom = async function(c) {
            const {
                title = "",
                html = ""
            } = c;
            const { value: formValues } = await Swal.fire({
                title: title,
                html: html,
                focusConfirm: false,
                preConfirm: () => {
                    return [
                        document.getElementById('swal-input1').value,
                        document.getElementById('swal-input2').value
                    ]
                }
            })

            if (formValues) {
                Swal.fire(JSON.stringify(formValues))
            }
        }
        return {
            toast: toast,
            success: success,
            error: error,
            custom: custom
        }
    }
    document.getElementById("check-availability").addEventListener("click", function(){
        console.log('hey')
        let html = '<input id="swal-input1" class="swal2-input">' +
                   '<input id="swal-input2" class="swal2-input">'
        attention.custom({title: "Make Reservation", html: html})
    })
})()