new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        username: null, // Our username
        password: null, // Our Password
        joined: false // True if user successfully logs in
    },

    mounted: function() {
        var self = this;
        if(!localStorage.getItem("loggedin")) {
            return;
        }

        self.joined = true;
        self.username = localStorage.getItem("loggedin");

        this.populate(self);
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {

            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">' + msg.username + '</div>' + msg.message + '<br/>';

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },

    methods: {
        populate: function(context) {
            console.log("populate field");

            // make a get to retrieve the past 20 messages
            fetch('http://localhost:8080/messages?page=0&size=20')
                .then(function(response){
                    response.text().then(function(text) {
                        var messages = JSON.parse(text);

                        for(i in messages.data) {
                            context.chatContent += '<div class="chip">' +
                                messages.data[i].username + '</div>' + messages.data[i].message + '<br/>';

                            var element = document.getElementById('chat-messages');
                            element.scrollTop = element.scrollHeight;
                        }
                    });
                });
        },

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

        logout: function() {
            this.joined = false;
            localStorage.removeItem("loggedin");
            this.chatContent = "";
            this.username = "";
            this.password = "";
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
                this.joined = (res.status == 200);
                localStorage.setItem('loggedin', this.username);

                if(this.joined) {
                    this.populate(this);
                }

                if(!this.joined) {
                    localStorage.removeItem("loggedin");
                    Materialize.toast("Invalid username or password", 2000);
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
                this.joined = (res.status == 201);
                localStorage.setItem('loggedin', this.username);

                if(this.joined) {
                    this.populate(this);
                }

                if(!this.joined) {
                    localStorage.removeItem("loggedin");
                    Materialize.toast("This username is already exists.", 2000);
                }
            });
        }
    }
});
