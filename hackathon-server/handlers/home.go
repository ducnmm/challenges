package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Hackathon Server</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .section { margin: 20px 0; padding: 15px; border: 1px solid #ddd; border-radius: 5px; }
        input, button { margin: 5px; padding: 8px; }
        button { background: #007cba; color: white; border: none; border-radius: 3px; cursor: pointer; }
        button:hover { background: #005a87; }
        #result { background: #f5f5f5; padding: 10px; border-radius: 3px; white-space: pre-wrap; }
        .token { background: #e8f5e8; padding: 10px; border-radius: 3px; margin: 10px 0; word-break: break-all; }
    </style>
</head>
<body>
    <h1>Hackathon Server - JWT Auth & File Upload</h1>
    
    <div class="section">
        <h2>Authentication</h2>
        
        <h3>Register</h3>
        <div>
            <input type="text" id="reg-username" placeholder="Username" />
            <input type="password" id="reg-password" placeholder="Password" />
            <button onclick="register()">Register</button>
        </div>
        
        <h3>Login</h3>
        <div>
            <input type="text" id="login-username" placeholder="Username" />
            <input type="password" id="login-password" placeholder="Password" />
            <button onclick="login()">Login</button>
        </div>
        
        <div id="token-display"></div>
    </div>
    
    <div class="section">
        <h2>File Upload</h2>
        <p><strong>Requirements:</strong> Must be logged in, file must be an image, max 8MB</p>
        <div>
            <input type="file" id="fileInput" accept="image/*" />
            <button onclick="uploadFile()">Upload</button>
        </div>
    </div>
    
    <div class="section">
        <h2>Response</h2>
        <div id="result">Ready to test...</div>
    </div>

    <script>
        let currentToken = '';
        
        async function register() {
            const username = document.getElementById('reg-username').value;
            const password = document.getElementById('reg-password').value;
            
            if (!username || !password) {
                alert('Please enter username and password');
                return;
            }
            
            try {
                const response = await fetch('/register', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, password })
                });
                const data = await response.json();
                document.getElementById('result').textContent = JSON.stringify(data, null, 2);
            } catch (error) {
                document.getElementById('result').textContent = 'Error: ' + error.message;
            }
        }
        
        async function login() {
            const username = document.getElementById('login-username').value;
            const password = document.getElementById('login-password').value;
            
            if (!username || !password) {
                alert('Please enter username and password');
                return;
            }
            
            try {
                const response = await fetch('/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, password })
                });
                const data = await response.json();
                
                if (data.success && data.data && data.data.token) {
                    currentToken = data.data.token;
                    document.getElementById('token-display').innerHTML = 
                        '<div class="token"><strong>JWT Token:</strong><br>' + currentToken + '</div>';
                }
                
                document.getElementById('result').textContent = JSON.stringify(data, null, 2);
            } catch (error) {
                document.getElementById('result').textContent = 'Error: ' + error.message;
            }
        }
        
        async function uploadFile() {
            if (!currentToken) {
                alert('Please login first to get a JWT token');
                return;
            }
            
            const fileInput = document.getElementById('fileInput');
            const file = fileInput.files[0];
            
            if (!file) {
                alert('Please select a file');
                return;
            }
            
            const formData = new FormData();
            formData.append('data', file);
            
            try {
                const response = await fetch('/upload', {
                    method: 'POST',
                    headers: { 'Authorization': 'Bearer ' + currentToken },
                    body: formData
                });
                const data = await response.json();
                document.getElementById('result').textContent = JSON.stringify(data, null, 2);
            } catch (error) {
                document.getElementById('result').textContent = 'Error: ' + error.message;
            }
        }
    </script>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}