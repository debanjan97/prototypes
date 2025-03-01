from flask import Flask
import os

app = Flask(__name__)

SERVER_ID = os.environ.get('SERVER_ID', 'unknown')

@app.route('/health')
def health():
    return f"I am server {SERVER_ID}"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=int(os.environ.get('PORT', 5000))) 