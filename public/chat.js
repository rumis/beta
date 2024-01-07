
document.addEventListener('DOMContentLoaded', function () {

    console.log('chat.js loaded');

    // clear chat room
    document.getElementById("btn_clear").addEventListener("click", function () {
        // ClearChatPanel();
        HideChatPanel();
        ShowWelcomePanel();
    });

    // send a message
    document.getElementById("btn_send").addEventListener("click", function () {
        let prompt = document.getElementById("prompt-textarea").value
        if (prompt == null || prompt == '') {
            return;
        }
        HideWelcomePanel();
        ShowChatPanel();

        StartChat(prompt, function (result) {
            console.log('callback result: ' + result);
        })
    })

});

// Show the welcome panel
function ShowWelcomePanel() {
    console.log('set welcome panel display to block');
    document.getElementById("chat_welcome").style.display = "flex";
}
// Hide the welcome panel
function HideWelcomePanel() {
    console.log('set welcome panel display to none');
    document.getElementById("chat_welcome").style.display = "none";
}
// Show the chat panel
function ShowChatPanel() {
    console.log('set chat panel display to block');
    document.getElementById("chat_room").style.display = "block";
}
// Hide the chat panel
function HideChatPanel() {
    console.log('set chat panel display to none');
    document.getElementById("chat_room").style.display = "none";
}
// clear the chat panel
function ClearChatPanel() {
    removeAllChildNodes(document.getElementById("chat_room"));
}

function ChatPanelAppendQuestion(question) {
}

// start a chat
function StartChat(prompt, callback) {
    fetch('/chat', {
        method: 'POST',
        body: JSON.stringify({
            prompt: prompt
        })
    }).then(response => {
        if (!response.ok) {
            throw new Error('network response is not ok');
        }
        return response.json(); // 或者 response.text() 如果期望的是文本内容
    }).then(result => {
        console.log('Success:', result);
        if (callback != null) {
            callback(result);
        }
    }).catch(error => {
        console.error('Error:', error);
    });
}


// remove all child nodes from a parent node
function removeAllChildNodes(parent) {
    while (parent.firstChild) {
        parent.removeChild(parent.firstChild);
    }
}


