new Vue({
    el: '#app',

    data: {
        apiUrl: "http://localhost:8080/",
        loginPath: "login",
        createUserPath: "user",
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        username: null, // Our username
        password: null, // Our Password
        joined: false // True if user successfully logs in
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {

            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">' + msg.username + '</div>'
                + emojione.toImage(msg.message) + '<br/>';

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                            username: this.username,
                            message: $('<p>').html(this.newMsg).text() // Strip out html
                        }
                    ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        login: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return;
            }

            if (!this.password) {
                Materialize.toast('You must enter a password', 2000);
                return;
            }

            fetch('http://localhost:8080/login', {
                method: 'POST',
                body:JSON.stringify(
                    {
                        username:this.username,
                        password:this.password
                    })
            }).then((res)=>{
                console.log("res");
                console.log(res.headers.get("token"));

                localStorage.setItem('token', res.body.token);

                this.joined = (res.status == 200);
                if(!this.joined) {
                    Materialize.toast("Invalid username or password");
                }
            });
        },

        create: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return;
            }

            if (!this.password) {
                Materialize.toast('You must enter a password', 2000);
                return;
            }

            fetch('http://localhost:8080/user', {
                method: 'POST',
                body:JSON.stringify(
                    {
                        username: this.username,
                        password: this.password
                    })
            }).then((res)=>{
                console.log("res");
                console.log(res.body.getReader().read());

                localStorage.setItem('token', res.body.token);
                this.joined = (res.status == 201);
                if(!this.joined) {
                    Materialize.toast("User failed to create.");
                }
            });
        }
    }
});
