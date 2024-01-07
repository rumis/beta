package tpl

var TemplateIndex = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>chat</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }

        #chat-container {
            max-width: 600px;
            margin: 20px auto;
            border: 1px solid #ccc;
            border-radius: 8px;
            overflow: hidden;
        }

        #messages {
            list-style: none;
            padding: 20px;
            margin: 0;
            overflow-y: auto;
            max-height: 300px;
        }

        #input-container {
            display: flex;
            padding: 10px;
            background-color: #fff;
        }

        #message-input {
            flex-grow: 1;
            padding: 8px;
            border: 1px solid #ccc;
            border-radius: 4px;
            margin-right: 8px;
        }

        #send-button {
            padding: 8px;
            background-color: #4caf50;
            color: #fff;
            border: none;
            border-radius: 4px;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <div id="chat-container">
        <ul id="messages"></ul>
        <div id="input-container">
            <input type="text" id="message-input" placeholder="Type your message...">
            <button id="send-button" onclick="sendMessage()">Send</button>
        </div>
    </div>

    <script>
        function sendMessage() {
            const inputElement = document.getElementById("message-input");
            const message = inputElement.value.trim();

            if (message !== "") {
                appendMessage("You", message);
                inputElement.value = "";
            }
        }

        function appendMessage(sender, text) {
            const messagesElement = document.getElementById("messages");
            const li = document.createElement("li");
            li.textContent = '${sender}: ${text}';
            messagesElement.appendChild(li);

            // Scroll to the bottom of the message container
            messagesElement.scrollTop = messagesElement.scrollHeight;
        }
    </script>
</body>
</html>
`
